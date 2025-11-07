<template>
  <div class="comment-section">
    <div class="comment-header">
      <h3>评论 ({{ total }})</h3>
    </div>

    <!-- 发表评论表单 -->
    <el-card v-if="userStore.isLoggedIn" class="comment-form-card" shadow="never">
      <div class="comment-input-wrapper">
        <el-input
          v-model="newComment"
          type="textarea"
          :rows="4"
          placeholder="写下你的评论..."
          :disabled="submitting"
          ref="commentTextareaRef"
        />
        <div class="input-toolbar">
          <el-popover
            v-model:visible="showEmojiPicker"
            placement="top"
            :width="300"
            trigger="manual"
          >
            <template #reference>
              <el-button 
                text 
                size="small" 
                @click="showEmojiPicker = !showEmojiPicker"
                type="primary"
              >
                😊 表情
              </el-button>
            </template>
            <div class="emoji-picker">
              <div class="emoji-grid">
                <span 
                  v-for="emoji in emojis" 
                  :key="emoji"
                  class="emoji-item"
                  @click="insertEmoji(emoji, 'newComment')"
                >
                  {{ emoji }}
                </span>
              </div>
            </div>
          </el-popover>
        </div>
      </div>
      <div class="form-actions">
        <el-button 
          type="primary" 
          @click="submitComment"
          :loading="submitting"
          :disabled="!newComment.trim()"
        >
          发表评论
        </el-button>
      </div>
    </el-card>
    <el-alert v-else type="info" :closable="false" style="margin-bottom: 20px">
      <template #default>
        请先 <el-link type="primary" @click="router.push('/login')">登录</el-link> 后再发表评论
      </template>
    </el-alert>

    <!-- 评论列表 -->
    <div class="comments-list" v-loading="loading">
      <el-empty v-if="comments.length === 0 && !loading" description="暂无评论，快来抢沙发吧！" />
      
      <div 
        v-for="comment in comments" 
        :key="comment.id" 
        :id="`comment-${comment.id}`"
        class="comment-item"
        :class="{ 'is-highlighted': highlightCommentId === comment.id }"
      >
        <div class="comment-avatar">
          <UserAvatar
            :avatar="comment.user?.avatar"
            :username="comment.user?.username"
            :nickname="comment.user?.nickname"
            size="medium"
          />
        </div>
        <div class="comment-content">
          <div class="comment-meta">
            <span class="username">{{ comment.user?.nickname || comment.user?.username }}</span>
            <span class="date">{{ formatDate(comment.created_at) }}</span>
          </div>
          <div class="comment-text" v-if="editingCommentId !== comment.id" v-html="renderContent(comment.content)"></div>
          <div v-else class="comment-edit">
            <div class="comment-input-wrapper">
              <el-input
                v-model="editContent"
                type="textarea"
                :rows="3"
                ref="editTextareaRef"
              />
              <div class="input-toolbar">
                <el-popover
                  v-model:visible="showEditEmojiPicker"
                  placement="top"
                  :width="300"
                  trigger="manual"
                >
                  <template #reference>
                    <el-button 
                      text 
                      size="small" 
                      @click="showEditEmojiPicker = !showEditEmojiPicker"
                      type="primary"
                    >
                      😊 表情
                    </el-button>
                  </template>
                  <div class="emoji-picker">
                    <div class="emoji-grid">
                      <span 
                        v-for="emoji in emojis" 
                        :key="emoji"
                        class="emoji-item"
                        @click="insertEmoji(emoji, 'editContent')"
                      >
                        {{ emoji }}
                      </span>
                    </div>
                  </div>
                </el-popover>
              </div>
            </div>
            <div class="edit-actions">
              <el-button size="small" @click="cancelEdit">取消</el-button>
              <el-button size="small" type="primary" @click="saveEdit(comment.id)">保存</el-button>
            </div>
          </div>
          <div class="comment-actions">
            <el-button 
              text 
              size="small"
              @click="showReplyForm(comment.id)"
              v-if="userStore.isLoggedIn"
            >
              <el-icon><ChatLineRound /></el-icon>
              回复
            </el-button>
            <el-button 
              text 
              size="small"
              @click="startEdit(comment)"
              v-if="isMyComment(comment)"
            >
              <el-icon><Edit /></el-icon>
              编辑
            </el-button>
            <el-button 
              text 
              size="small"
              type="danger"
              @click="deleteComment(comment.id)"
              v-if="isMyComment(comment)"
            >
              <el-icon><Delete /></el-icon>
              删除
            </el-button>
          </div>

          <!-- 回复表单 -->
          <div v-if="replyingTo === comment.id" class="reply-form">
            <div class="comment-input-wrapper">
              <el-input
                v-model="replyContent"
                type="textarea"
                :rows="3"
                :placeholder="`回复 ${comment.user?.nickname || comment.user?.username}...`"
                ref="replyTextareaRef"
              />
              <div class="input-toolbar">
                <el-popover
                  v-model:visible="showReplyEmojiPicker"
                  placement="top"
                  :width="300"
                  trigger="manual"
                >
                  <template #reference>
                    <el-button 
                      text 
                      size="small" 
                      @click="showReplyEmojiPicker = !showReplyEmojiPicker"
                      type="primary"
                    >
                      😊 表情
                    </el-button>
                  </template>
                  <div class="emoji-picker">
                    <div class="emoji-grid">
                      <span 
                        v-for="emoji in emojis" 
                        :key="emoji"
                        class="emoji-item"
                        @click="insertEmoji(emoji, 'replyContent')"
                      >
                        {{ emoji }}
                      </span>
                    </div>
                  </div>
                </el-popover>
              </div>
            </div>
            <div class="form-actions">
              <el-button size="small" @click="cancelReply">取消</el-button>
              <el-button 
                size="small" 
                type="primary" 
                @click="submitReply(comment.id)"
                :loading="submitting"
                :disabled="!replyContent.trim()"
              >
                发表回复
              </el-button>
            </div>
          </div>

          <!-- 回复列表 -->
          <div v-if="comment.replies && comment.replies.length > 0" class="replies-list">
            <div 
              v-for="reply in comment.replies" 
              :key="reply.id" 
              :id="`comment-${reply.id}`"
              class="reply-item"
              :class="{ 'is-highlighted': highlightCommentId === reply.id }"
            >
              <div class="comment-avatar">
                <UserAvatar
                  :avatar="reply.user?.avatar"
                  :username="reply.user?.username"
                  :nickname="reply.user?.nickname"
                  size="small"
                />
              </div>
              <div class="comment-content">
                <div class="comment-meta">
                  <span class="username">{{ reply.user?.nickname || reply.user?.username }}</span>
                  <span class="reply-indicator">回复了</span>
                  <span class="reply-to-user">{{ getReplyToUser(reply.content, comment.user) }}</span>
                  <span class="date">{{ formatDate(reply.created_at) }}</span>
                </div>
                <div class="comment-text" v-html="renderContent(reply.content)"></div>
                <div class="comment-actions">
                  <el-button 
                    text 
                    size="small"
                    @click="showReplyForm(reply.id)"
                    v-if="userStore.isLoggedIn"
                  >
                    <el-icon><ChatLineRound /></el-icon>
                    回复
                  </el-button>
                  <el-button 
                    text 
                    size="small"
                    type="danger"
                    @click="deleteComment(reply.id)"
                    v-if="isMyComment(reply)"
                  >
                    <el-icon><Delete /></el-icon>
                    删除
                  </el-button>
                </div>
                
                <!-- 回复的回复表单 -->
                <div v-if="replyingTo === reply.id" class="reply-form">
                  <div class="comment-input-wrapper">
                    <el-input
                      v-model="replyContent"
                      type="textarea"
                      :rows="3"
                      :placeholder="`回复 ${reply.user?.nickname || reply.user?.username}...`"
                      ref="replyTextareaRef"
                    />
                    <div class="input-toolbar">
                      <el-popover
                        v-model:visible="showReplyEmojiPicker"
                        placement="top"
                        :width="300"
                        trigger="manual"
                      >
                        <template #reference>
                          <el-button 
                            text 
                            size="small" 
                            @click="showReplyEmojiPicker = !showReplyEmojiPicker"
                            type="primary"
                          >
                            😊 表情
                          </el-button>
                        </template>
                        <div class="emoji-picker">
                          <div class="emoji-grid">
                            <span 
                              v-for="emoji in emojis" 
                              :key="emoji"
                              class="emoji-item"
                              @click="insertEmoji(emoji, 'replyContent')"
                            >
                              {{ emoji }}
                            </span>
                          </div>
                        </div>
                      </el-popover>
                    </div>
                  </div>
                  <div class="form-actions">
                    <el-button size="small" @click="cancelReply">取消</el-button>
                    <el-button 
                      size="small" 
                      type="primary" 
                      @click="submitReply(reply.id)"
                      :loading="submitting"
                      :disabled="!replyContent.trim()"
                    >
                      发表回复
                    </el-button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ChatLineRound, Edit, Delete } from '@element-plus/icons-vue'
import { commentAPI } from '@/api/comment'
import { useUserStore } from '@/stores/user'
import UserAvatar from '@/components/UserAvatar.vue'

const props = defineProps({
  articleId: {
    type: Number,
    required: true
  },
  highlightCommentId: {
    type: Number,
    default: null
  }
})

const router = useRouter()
const userStore = useUserStore()

const loading = ref(false)
const submitting = ref(false)
const comments = ref([])
const total = ref(0)
const newComment = ref('')
const replyingTo = ref(null)
const replyContent = ref('')
const editingCommentId = ref(null)
const editContent = ref('')
const showEmojiPicker = ref(false)
const showReplyEmojiPicker = ref(false)
const showEditEmojiPicker = ref(false)
const commentTextareaRef = ref(null)
const replyTextareaRef = ref(null)

// 表情包列表
const emojis = [
  '😀', '😃', '😄', '😁', '😆', '😅', '🤣', '😂', '🙂', '🙃',
  '😉', '😊', '😇', '🥰', '😍', '🤩', '😘', '😗', '😚', '😙',
  '😋', '😛', '😜', '🤪', '😝', '🤑', '🤗', '🤭', '🤫', '🤔',
  '🤐', '🤨', '😐', '😑', '😶', '😏', '😒', '🙄', '😬', '🤥',
  '😌', '😔', '😪', '🤤', '😴', '😷', '🤒', '🤕', '🤢', '🤮',
  '🤧', '🥵', '🥶', '😶‍🌫️', '😵', '😵‍💫', '🤯', '🤠', '🥳', '🥸',
  '😎', '🤓', '🧐', '😕', '😟', '🙁', '☹️', '😮', '😯', '😲',
  '😳', '🥺', '😦', '😧', '😨', '😰', '😥', '😢', '😭', '😱',
  '😖', '😣', '😞', '😓', '😩', '😫', '🥱', '😤', '😡', '😠',
  '🤬', '😈', '👿', '💀', '☠️', '💩', '🤡', '👹', '👺', '👻',
  '👽', '👾', '🤖', '😺', '😸', '😹', '😻', '😼', '😽', '🙀',
  '😿', '😾', '💋', '💌', '💘', '💝', '💖', '💗', '💓', '💞',
  '💕', '💟', '❣️', '💔', '❤️', '🧡', '💛', '💚', '💙', '💜',
  '🤎', '🖤', '🤍', '💯', '💢', '💥', '💫', '💦', '💨', '🕳️',
  '💣', '💬', '👁️‍🗨️', '🗨️', '🗯️', '💭', '💤', '👋', '🤚', '🖐️',
  '✋', '🖖', '👌', '🤌', '🤏', '✌️', '🤞', '🤟', '🤘', '🤙',
  '👈', '👉', '👆', '🖕', '👇', '☝️', '👍', '👎', '✊', '👊',
  '🤛', '🤜', '👏', '🙌', '👐', '🤲', '🤝', '🙏', '✍️', '💅',
  '🤳', '💪', '🦾', '🦿', '🦵', '🦶', '👂', '🦻', '👃', '👶',
  '👧', '🧒', '👦', '👩', '🧑', '👨', '👩‍🦱', '🧑‍🦱', '👨‍🦱', '👩‍🦰',
  '🧑‍🦰', '👨‍🦰', '👱‍♀️', '👱', '👱‍♂️', '👩‍🦳', '🧑‍🦳', '👨‍🦳', '👩‍🦲', '🧑‍🦲',
  '👨‍🦲', '🧔', '👵', '🧓', '👴', '👲', '👳‍♀️', '👳', '👳‍♂️', '🧕',
  '👮‍♀️', '👮', '👮‍♂️', '👷‍♀️', '👷', '👷‍♂️', '💂‍♀️', '💂', '💂‍♂️', '🕵️‍♀️',
  '🕵️', '🕵️‍♂️', '👩‍⚕️', '🧑‍⚕️', '👨‍⚕️', '👩‍🌾', '🧑‍🌾', '👨‍🌾', '👩‍🍳', '🧑‍🍳',
  '👨‍🍳', '👩‍🎓', '🧑‍🎓', '👨‍🎓', '👩‍🎤', '🧑‍🎤', '👨‍🎤', '👩‍🏫', '🧑‍🏫', '👨‍🏫',
  '👩‍🏭', '🧑‍🏭', '👨‍🏭', '👩‍💻', '🧑‍💻', '👨‍💻', '👩‍💼', '🧑‍💼', '👨‍💼', '👩‍🔧',
  '🧑‍🔧', '👨‍🔧', '👩‍🔬', '🧑‍🔬', '👨‍🔬', '👩‍🎨', '🧑‍🎨', '👨‍🎨', '👩‍🚒', '🧑‍🚒',
  '👨‍🚒', '👩‍✈️', '🧑‍✈️', '👨‍✈️', '👩‍🚀', '🧑‍🚀', '👨‍🚀', '👩‍⚖️', '🧑‍⚖️', '👨‍⚖️',
  '👰', '🤵', '👸', '🤴', '🦸‍♀️', '🦸', '🦸‍♂️', '🦹‍♀️', '🦹', '🦹‍♂️',
  '🤶', '🎅', '🧙‍♀️', '🧙', '🧙‍♂️', '🧝‍♀️', '🧝', '🧝‍♂️', '🧛‍♀️', '🧛',
  '🧛‍♂️', '🧟‍♀️', '🧟', '🧟‍♂️', '🧞‍♀️', '🧞', '🧞‍♂️', '🧜‍♀️', '🧜', '🧜‍♂️',
  '🧚‍♀️', '🧚', '🧚‍♂️', '👼', '🤰', '🤱', '👩‍🍼', '🧑‍🍼', '👨‍🍼', '🙇‍♀️',
  '🙇', '🙇‍♂️', '💁‍♀️', '💁', '💁‍♂️', '🙅‍♀️', '🙅', '🙅‍♂️', '🙆‍♀️', '🙆',
  '🙆‍♂️', '🙋‍♀️', '🙋', '🙋‍♂️', '🧏‍♀️', '🧏', '🧏‍♂️', '🤦‍♀️', '🤦', '🤦‍♂️',
  '🤷‍♀️', '🤷', '🤷‍♂️', '🧑‍🤝‍🧑', '👭', '👫', '👬', '💏', '💑', '👪',
  '👨‍👩‍👧', '👨‍👩‍👧‍👦', '👨‍👩‍👦‍👦', '👨‍👩‍👧‍👧', '👩‍👩‍👦', '👩‍👩‍👧', '👩‍👩‍👧‍👦', '👩‍👩‍👦‍👦', '👩‍👩‍👧‍👧', '👨‍👨‍👦',
  '👨‍👨‍👧', '👨‍👨‍👧‍👦', '👨‍👨‍👦‍👦', '👨‍👨‍👧‍👧', '👩‍👦', '👩‍👧', '👩‍👧‍👦', '👩‍👦‍👦', '👩‍👧‍👧', '👨‍👦',
  '👨‍👧', '👨‍👧‍👦', '👨‍👦‍👦', '👨‍👧‍👧', '🗣️', '👤', '👥', '👣', '🌍', '🌎',
  '🌏', '🗺️', '🧭', '🏔️', '⛰️', '🌋', '🗻', '🏕️', '🏖️', '🏜️',
  '🏝️', '🏞️', '🏟️', '🏛️', '🏗️', '🧱', '🏘️', '🏚️', '🏠', '🏡',
  '🏢', '🏣', '🏤', '🏥', '🏦', '🏨', '🏩', '🏪', '🏫', '🏬',
  '🏭', '🏯', '🏰', '💒', '🗼', '🗽', '⛪', '🕌', '🛕', '🕍',
  '⛩️', '🕋', '⛲', '⛺', '🌁', '🌃', '🏙️', '🌄', '🌅', '🌆',
  '🌇', '🌉', '♨️', '🎠', '🎡', '🎢', '💈', '🎪', '🚂', '🚃',
  '🚄', '🚅', '🚆', '🚇', '🚈', '🚉', '🚊', '🚝', '🚞', '🚋',
  '🚌', '🚍', '🚎', '🚐', '🚑', '🚒', '🚓', '🚔', '🚕', '🚖',
  '🚗', '🚘', '🚙', '🚚', '🚛', '🚜', '🏎️', '🏍️', '🛵', '🚲',
  '🛴', '🛹', '🛼', '🚁', '🛸', '✈️', '🛫', '🛬', '🪂', '💺',
  '🚢', '⛵', '🛥️', '🛳️', '⛴️', '🚤', '🛶', '🚁', '🚟', '🚠',
  '🚡', '🛰️', '🚀', '🛸', '🎆', '🎇', '🎈', '🎉', '🎊', '🎋',
  '🎍', '🎎', '🎏', '🎐', '🎑', '🧧', '🎀', '🎁', '🎗️', '🎟️',
  '🎫', '🎖️', '🏆', '🏅', '🥇', '🥈', '🥉', '⚽', '⚾', '🥎',
  '🏀', '🏐', '🏈', '🏉', '🎾', '🥏', '🎳', '🏏', '🏑', '🏒',
  '🥍', '🏓', '🏸', '🥊', '🥋', '🥅', '⛳', '⛸️', '🎣', '🤿',
  '🎽', '🎿', '🛷', '🥌', '🎯', '🎱', '🪀', '🎮', '🕹️', '🎰',
  '🎲', '🧩', '🧸', '♟️', '🃏', '🀄', '🎴', '🎭', '🖼️', '🎨',
  '🧵', '🧶', '👓', '🕶️', '🥽', '🥼', '🦺', '👔', '👕', '👖',
  '🧣', '🧤', '🧥', '🧦', '👗', '👘', '🥻', '🩱', '🩲', '🩳',
  '👙', '👚', '👛', '👜', '👝', '🛍️', '🎒', '👞', '👟', '🥾',
  '🥿', '👠', '👡', '🩰', '👢', '🪖', '⛑️', '🎩', '🎓', '🧢',
  '👑', '💎', '⚖️', '🔧', '🔨', '⚒️', '🛠️', '⛏️', '🔩', '⚙️',
  '🧰', '🧲', '🪚', '🔫', '💣', '🧨', '🪓', '🔪', '🗡️', '⚔️',
  '🛡️', '🚬', '⚰️', '🪦', '⚱️', '🏺', '🔮', '📿', '🧿', '💈',
  '⚗️', '🔭', '🔬', '🕳️', '🩹', '🩺', '💊', '💉', '🩸', '🧬',
  '🦠', '🧫', '🧪', '🌡️', '🧹', '🪠', '🧺', '🧻', '🚽', '🚿',
  '🛁', '🛀', '🧼', '🪥', '🪒', '🧽', '🧴', '🛎️', '🔑', '🗝️',
  '🚪', '🪑', '🛋️', '🛏️', '🛌', '🧸', '🪆', '🖼️', '🪞', '🪟',
  '🛍️', '🛒', '🎁', '🎈', '🎉', '🎊', '🎀', '🎗️', '🏆', '🥇',
  '🥈', '🥉', '⚽', '🏀', '⚾', '🎾', '🏐', '🏉', '🎱', '🏓',
  '🏸', '🥅', '⛳', '🏌️', '🏌️‍♂️', '🏌️‍♀️', '🏄', '🏄‍♂️', '🏄‍♀️', '🏊',
  '🏊‍♂️', '🏊‍♀️', '⛷️', '🏂', '🪂', '🏋️', '🏋️‍♂️', '🏋️‍♀️', '🚴', '🚴‍♂️',
  '🚴‍♀️', '🚵', '🚵‍♂️', '🚵‍♀️', '🤸', '🤸‍♂️', '🤸‍♀️', '🤽', '🤽‍♂️', '🤽‍♀️',
  '🤾', '🤾‍♂️', '🤾‍♀️', '🤹', '🤹‍♂️', '🤹‍♀️', '🧘', '🧘‍♂️', '🧘‍♀️', '🛀',
  '🛌', '👭', '👫', '👬', '💏', '💑', '👪', '🗣️', '👤', '👥',
  '👣', '🐶', '🐱', '🐭', '🐹', '🐰', '🦊', '🐻', '🐼', '🐨',
  '🐯', '🦁', '🐮', '🐷', '🐽', '🐸', '🐵', '🙈', '🙉', '🙊',
  '🐒', '🐔', '🐧', '🐦', '🐤', '🐣', '🐥', '🦆', '🦅', '🦉',
  '🦇', '🐺', '🐗', '🐴', '🦄', '🐝', '🐛', '🦋', '🐌', '🐞',
  '🐜', '🪰', '🪱', '🦟', '🦗', '🕷️', '🦂', '🐢', '🐍', '🦎',
  '🦖', '🦕', '🐙', '🦑', '🦐', '🦞', '🦀', '🐡', '🐠', '🐟',
  '🐬', '🐳', '🐋', '🦈', '🐊', '🐅', '🐆', '🦓', '🦍', '🦧',
  '🐘', '🦛', '🦏', '🐪', '🐫', '🦒', '🦘', '🦬', '🐃', '🐂',
  '🐄', '🐎', '🐖', '🐏', '🐑', '🦙', '🐐', '🦌', '🐕', '🐩',
  '🦮', '🐕‍🦺', '🐈', '🐈‍⬛', '🪶', '🐓', '🦃', '🦤', '🦚', '🦜',
  '🦢', '🦩', '🕊️', '🐇', '🦝', '🦨', '🦡', '🦫', '🦦', '🦥',
  '🐁', '🐀', '🐿️', '🦔', '🐾', '🐉', '🐲', '🌵', '🎄', '🌲',
  '🌳', '🌴', '🪵', '🌱', '🌿', '☘️', '🍀', '🎍', '🪴', '🎋',
  '🍃', '🍂', '🍁', '🍄', '🐚', '🪨', '🌾', '💐', '🌷', '🌹',
  '🥀', '🌺', '🌻', '🌼', '🌷', '🌱', '🌿', '🍀', '☘️', '🍃',
  '🍂', '🍁', '🍄', '🌰', '🪵', '🦀', '🦞', '🦐', '🦑', '🦪',
  '🍇', '🍈', '🍉', '🍊', '🍋', '🍌', '🍍', '🥭', '🍎', '🍏',
  '🍐', '🍑', '🍒', '🍓', '🫐', '🥝', '🍅', '🥥', '🥑', '🍆',
  '🥔', '🥕', '🌽', '🌶️', '🫑', '🥒', '🥬', '🥦', '🧄', '🧅',
  '🍄', '🥜', '🌰', '🍞', '🥐', '🥖', '🫓', '🥨', '🥯', '🥞',
  '🧇', '🍳', '🥚', '🧀', '🥓', '🥩', '🍗', '🍖', '🦴', '🌭',
  '🍔', '🍟', '🍕', '🫔', '🥪', '🥙', '🧆', '🌮', '🌯', '🫔',
  '🥗', '🥘', '🥫', '🍝', '🍜', '🍲', '🍛', '🍣', '🍱', '🥟',
  '🦪', '🍤', '🍙', '🍚', '🍘', '🍥', '🥠', '🥮', '🍢', '🍡',
  '🍧', '🍨', '🍦', '🥧', '🧁', '🍰', '🎂', '🍮', '🍭', '🍬',
  '🍫', '🍿', '🍩', '🍪', '🌰', '🥜', '🍯', '🥛', '🍼', '🫖',
  '☕', '🍵', '🧃', '🥤', '🧋', '🍶', '🍺', '🍻', '🥂', '🍷',
  '🥃', '🍸', '🍹', '🧉', '🍾', '🧊', '🥄', '🍴', '🍽️', '🥣',
  '🥡', '🥢', '🪣', '🧂', '⚽', '🏀', '⚾', '🥎', '🏐', '🏈',
  '🎾', '🏐', '🏉', '🎱', '🏓', '🏸', '🥅', '⛳', '🏌️', '🏄',
  '🏊', '⛷️', '🏂', '🪂', '🏋️', '🚴', '🚵', '🤸', '🤽', '🤾',
  '🤹', '🧘', '🛀', '🛌', '👭', '👫', '👬', '💏', '💑', '👪',
  '🗣️', '👤', '👥', '👣', '🌍', '🌎', '🌏', '🗺️', '🧭', '🏔️',
  '⛰️', '🌋', '🗻', '🏕️', '🏖️', '🏜️', '🏝️', '🏞️', '🏟️', '🏛️',
  '🏗️', '🧱', '🏘️', '🏚️', '🏠', '🏡', '🏢', '🏣', '🏤', '🏥',
  '🏦', '🏨', '🏩', '🏪', '🏫', '🏬', '🏭', '🏯', '🏰', '💒',
  '🗼', '🗽', '⛪', '🕌', '🛕', '🕍', '⛩️', '🕋', '⛲', '⛺',
  '🌁', '🌃', '🏙️', '🌄', '🌅', '🌆', '🌇', '🌉', '♨️', '🎠',
  '🎡', '🎢', '💈', '🎪', '🚂', '🚃', '🚄', '🚅', '🚆', '🚇',
  '🚈', '🚉', '🚊', '🚝', '🚞', '🚋', '🚌', '🚍', '🚎', '🚐',
  '🚑', '🚒', '🚓', '🚔', '🚕', '🚖', '🚗', '🚘', '🚙', '🚚',
  '🚛', '🚜', '🏎️', '🏍️', '🛵', '🚲', '🛴', '🛹', '🛼', '🚁',
  '🛸', '✈️', '🛫', '🛬', '🪂', '💺', '🚢', '⛵', '🛥️', '🛳️',
  '⛴️', '🚤', '🛶', '🚁', '🚟', '🚠', '🚡', '🛰️', '🚀', '🛸',
  '🎆', '🎇', '🎈', '🎉', '🎊', '🎋', '🎍', '🎎', '🎏', '🎐',
  '🎑', '🧧', '🎀', '🎁', '🎗️', '🎟️', '🎫', '🎖️', '🏆', '🏅',
  '🥇', '🥈', '🥉', '⚽', '⚾', '🥎', '🏀', '🏐', '🏈', '🏉',
  '🎾', '🥏', '🎳', '🏏', '🏑', '🏒', '🥍', '🏓', '🏸', '🥊',
  '🥋', '🥅', '⛳', '⛸️', '🎣', '🤿', '🎽', '🎿', '🛷', '🥌',
  '🎯', '🎱', '🪀', '🎮', '🕹️', '🎰', '🎲', '🧩', '🧸', '♟️',
  '🃏', '🀄', '🎴', '🎭', '🖼️', '🎨', '🧵', '🧶', '👓', '🕶️',
  '🥽', '🥼', '🦺', '👔', '👕', '👖', '🧣', '🧤', '🧥', '🧦',
  '👗', '👘', '🥻', '🩱', '🩲', '🩳', '👙', '👚', '👛', '👜',
  '👝', '🛍️', '🎒', '👞', '👟', '🥾', '🥿', '👠', '👡', '🩰',
  '👢', '🪖', '⛑️', '🎩', '🎓', '🧢', '👑', '💎', '⚖️', '🔧',
  '🔨', '⚒️', '🛠️', '⛏️', '🔩', '⚙️', '🧰', '🧲', '🪚', '🔫',
  '💣', '🧨', '🪓', '🔪', '🗡️', '⚔️', '🛡️', '🚬', '⚰️', '🪦',
  '⚱️', '🏺', '🔮', '📿', '🧿', '💈', '⚗️', '🔭', '🔬', '🕳️',
  '🩹', '🩺', '💊', '💉', '🩸', '🧬', '🦠', '🧫', '🧪', '🌡️',
  '🧹', '🪠', '🧺', '🧻', '🚽', '🚿', '🛁', '🛀', '🧼', '🪥',
  '🪒', '🧽', '🧴', '🛎️', '🔑', '🗝️', '🚪', '🪑', '🛋️', '🛏️',
  '🛌', '🧸', '🪆', '🖼️', '🪞', '🪟', '🛍️', '🛒', '🎁', '🎈',
  '🎉', '🎊', '🎀', '🎗️', '🏆', '🥇', '🥈', '🥉'
]

// 插入表情到输入框
const insertEmoji = (emoji, target) => {
  if (target === 'newComment') {
    const textarea = commentTextareaRef.value?.$el?.querySelector('textarea')
    if (textarea) {
      const start = textarea.selectionStart
      const end = textarea.selectionEnd
      newComment.value = newComment.value.substring(0, start) + emoji + newComment.value.substring(end)
      nextTick(() => {
        textarea.focus()
        const newPos = start + emoji.length
        textarea.setSelectionRange(newPos, newPos)
      })
    } else {
      newComment.value += emoji
    }
    showEmojiPicker.value = false
  } else if (target === 'replyContent') {
    const textarea = replyTextareaRef.value?.$el?.querySelector('textarea')
    if (textarea) {
      const start = textarea.selectionStart
      const end = textarea.selectionEnd
      replyContent.value = replyContent.value.substring(0, start) + emoji + replyContent.value.substring(end)
      nextTick(() => {
        textarea.focus()
        const newPos = start + emoji.length
        textarea.setSelectionRange(newPos, newPos)
      })
    } else {
      replyContent.value += emoji
    }
    showReplyEmojiPicker.value = false
  } else if (target === 'editContent') {
    const textarea = editTextareaRef.value?.$el?.querySelector('textarea')
    if (textarea) {
      const start = textarea.selectionStart
      const end = textarea.selectionEnd
      editContent.value = editContent.value.substring(0, start) + emoji + editContent.value.substring(end)
      nextTick(() => {
        textarea.focus()
        const newPos = start + emoji.length
        textarea.setSelectionRange(newPos, newPos)
      })
    } else {
      editContent.value += emoji
    }
    showEditEmojiPicker.value = false
  }
}

// 获取被回复者的用户名
const getReplyToUser = (content, topLevelCommentUser) => {
  if (!content) return topLevelCommentUser?.nickname || topLevelCommentUser?.username || ''
  
  // 从内容中提取 @用户名
  const match = content.match(/@(\S+)/)
  if (match && match[1]) {
    return match[1]
  }
  
  // 如果没有 @，返回顶层评论的用户
  return topLevelCommentUser?.nickname || topLevelCommentUser?.username || ''
}

// 渲染内容（支持表情和基本HTML）
const renderContent = (content) => {
  if (!content) return ''
  // 转义HTML，但保留换行
  let rendered = content
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;')
    .replace(/\n/g, '<br>')
  
  // 高亮显示 @用户名
  rendered = rendered.replace(/@(\S+)/g, '<span class="mention">@$1</span>')
  
  return rendered
}

const editTextareaRef = ref(null)

const loadComments = async () => {
  try {
    loading.value = true
    const res = await commentAPI.getArticleComments(props.articleId)
    comments.value = res.comments || []
    total.value = res.total || 0
  } catch (error) {
    console.error(error)
  } finally {
    loading.value = false
  }
}

const submitComment = async () => {
  try {
    submitting.value = true
    await commentAPI.createComment({
      content: newComment.value,
      article_id: props.articleId
    })
    ElMessage.success('评论发表成功')
    newComment.value = ''
    loadComments()
  } catch (error) {
    console.error(error)
  } finally {
    submitting.value = false
  }
}

const showReplyForm = (commentId) => {
  replyingTo.value = commentId
  replyContent.value = ''
}

const cancelReply = () => {
  replyingTo.value = null
  replyContent.value = ''
}

const submitReply = async (commentId) => {
  try {
    submitting.value = true
    
    // 找到顶层评论的ID和被回复者的用户名
    let topLevelCommentId = commentId
    let replyToUsername = ''
    
    // 遍历所有评论，检查commentId是否是某个回复
    for (const comment of comments.value) {
      if (comment.id === commentId) {
        // 这是顶层评论
        topLevelCommentId = commentId
        replyToUsername = comment.user?.nickname || comment.user?.username
        break
      }
      // 检查是否在回复列表中
      if (comment.replies && comment.replies.length > 0) {
        const foundReply = comment.replies.find(r => r.id === commentId)
        if (foundReply) {
          // 这是一个回复，使用顶层评论的ID
          topLevelCommentId = comment.id
          replyToUsername = foundReply.user?.nickname || foundReply.user?.username
          break
        }
      }
    }
    
    // 在回复内容前添加 @用户名
    const contentWithMention = `@${replyToUsername} ${replyContent.value}`
    
    const payload = {
      content: contentWithMention,
      article_id: props.articleId,
      parent_id: topLevelCommentId  // 始终使用顶层评论的ID
    }
    
    console.log('提交回复:', payload, '原始commentId:', commentId, '顶层ID:', topLevelCommentId)
    await commentAPI.createComment(payload)
    ElMessage.success('回复发表成功')
    cancelReply()
    await loadComments()
  } catch (error) {
    console.error('回复失败:', error)
    ElMessage.error(error.response?.data?.error || '回复失败，请重试')
  } finally {
    submitting.value = false
  }
}

const startEdit = (comment) => {
  editingCommentId.value = comment.id
  editContent.value = comment.content
}

const cancelEdit = () => {
  editingCommentId.value = null
  editContent.value = ''
}

const saveEdit = async (commentId) => {
  try {
    await commentAPI.updateComment(commentId, { content: editContent.value })
    ElMessage.success('评论更新成功')
    cancelEdit()
    loadComments()
  } catch (error) {
    console.error(error)
  }
}

const deleteComment = async (commentId) => {
  try {
    await ElMessageBox.confirm('确定要删除这条评论吗？', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'warning'
    })
    await commentAPI.deleteComment(commentId)
    ElMessage.success('评论删除成功')
    loadComments()
  } catch (error) {
    if (error !== 'cancel') {
      console.error(error)
    }
  }
}

const isMyComment = (comment) => {
  return userStore.user && comment.user_id === userStore.user.id
}

const formatDate = (dateStr) => {
  const date = new Date(dateStr)
  const now = new Date()
  const diff = now - date
  
  const minute = 60 * 1000
  const hour = 60 * minute
  const day = 24 * hour
  
  if (diff < minute) {
    return '刚刚'
  } else if (diff < hour) {
    return `${Math.floor(diff / minute)} 分钟前`
  } else if (diff < day) {
    return `${Math.floor(diff / hour)} 小时前`
  } else if (diff < 7 * day) {
    return `${Math.floor(diff / day)} 天前`
  } else {
    return date.toLocaleDateString('zh-CN')
  }
}

onMounted(() => {
  loadComments()
})
</script>

<style scoped>
.comment-section {
  margin-top: 30px;
}

.comment-header h3 {
  font-size: 24px;
  font-weight: 600;
  margin-bottom: 24px;
  color: #1d1d1f;
  letter-spacing: -0.011em;
}

.comment-form-card {
  margin-bottom: 24px;
  border-radius: 16px;
  border: 1px solid rgba(255, 182, 193, 0.2);
  background: rgba(255, 240, 248, 0.2);
}

.form-actions {
  margin-top: 16px;
  text-align: right;
}

.comments-list {
  margin-top: 24px;
}

.comment-item {
  display: flex;
  gap: 16px;
  padding: 24px 0;
  border-bottom: 1px solid rgba(255, 182, 193, 0.15);
  transition: background 0.2s;
  border-radius: 12px;
  padding-left: 12px;
  padding-right: 12px;
  margin-left: -12px;
  margin-right: -12px;
}

.comment-item:hover {
  background: rgba(255, 240, 248, 0.3);
}

.comment-item:last-child {
  border-bottom: none;
}

.comment-avatar {
  flex-shrink: 0;
}

.comment-content {
  flex: 1;
}

.comment-meta {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 10px;
}

.username {
  font-weight: 500;
  color: #1d1d1f;
  font-size: 15px;
}

.date {
  font-size: 13px;
  color: #86868b;
}

.reply-indicator {
  font-size: 13px;
  color: #86868b;
  font-weight: 400;
}

.reply-to-user {
  font-size: 13px;
  color: #ff6b9d;
  font-weight: 600;
}

.comment-text {
  color: #1d1d1f;
  line-height: 1.8;
  margin-bottom: 12px;
  white-space: pre-wrap;
  word-wrap: break-word;
  font-size: 15px;
}

.comment-actions {
  display: flex;
  gap: 12px;
}

.comment-edit {
  margin-bottom: 12px;
}

.edit-actions {
  margin-top: 12px;
  display: flex;
  gap: 8px;
  justify-content: flex-end;
}

.reply-form {
  margin-top: 16px;
  padding: 16px;
  background: rgba(255, 240, 248, 0.4);
  border-radius: 12px;
  border: 1px solid rgba(255, 182, 193, 0.2);
}

.replies-list {
  margin-top: 16px;
  padding-left: 24px;
  border-left: 2px solid rgba(255, 182, 193, 0.3);
}

.reply-item {
  display: flex;
  gap: 12px;
  padding: 16px 0;
}

.reply-item:first-child {
  padding-top: 0;
}

.reply-item:last-child {
  padding-bottom: 0;
}

:deep(.el-button--text) {
  color: #ff6b9d;
}

:deep(.el-button--text:hover) {
  color: #ff8fab;
  background: rgba(255, 240, 248, 0.3);
}

:deep(.el-button--text.is-disabled) {
  color: #c0c4cc;
}

:deep(.el-alert--info) {
  background-color: rgba(255, 240, 248, 0.5);
  border: 1px solid rgba(255, 182, 193, 0.3);
  border-radius: 12px;
}

:deep(.el-alert__title) {
  color: #1d1d1f;
}

:deep(.el-card) {
  border-radius: 16px;
  border: 1px solid rgba(255, 182, 193, 0.2);
  background: rgba(255, 255, 255, 0.9);
}

.comment-input-wrapper {
  position: relative;
}

.input-toolbar {
  margin-top: 8px;
  display: flex;
  gap: 8px;
}

.emoji-picker {
  max-height: 300px;
  overflow-y: auto;
  padding: 8px;
}

.emoji-grid {
  display: grid;
  grid-template-columns: repeat(10, 1fr);
  gap: 4px;
}

.emoji-item {
  font-size: 24px;
  cursor: pointer;
  padding: 4px;
  text-align: center;
  border-radius: 4px;
  transition: background 0.2s;
  user-select: none;
}

.emoji-item:hover {
  background: rgba(255, 182, 193, 0.3);
}

@media (max-width: 768px) {
  .emoji-grid {
    grid-template-columns: repeat(8, 1fr);
  }
  
  .emoji-item {
    font-size: 20px;
  }
}

/* @用户名 提及样式 */
:deep(.mention) {
  color: #ff6b9d;
  font-weight: 600;
  background: rgba(255, 182, 193, 0.15);
  padding: 2px 6px;
  border-radius: 4px;
  margin-right: 4px;
  transition: all 0.2s;
}

:deep(.mention:hover) {
  background: rgba(255, 182, 193, 0.25);
}

/* 高亮评论动画 */
.is-highlighted {
  animation: highlightFade 3s ease-in-out;
}

@keyframes highlightFade {
  0% {
    background: rgba(255, 107, 157, 0.2);
    box-shadow: 0 0 0 3px rgba(255, 107, 157, 0.3);
  }
  50% {
    background: rgba(255, 107, 157, 0.15);
    box-shadow: 0 0 0 3px rgba(255, 107, 157, 0.2);
  }
  100% {
    background: transparent;
    box-shadow: none;
  }
}

.comment-item.is-highlighted,
.reply-item.is-highlighted {
  border-radius: 12px;
  transition: all 0.3s ease;
}
</style>




