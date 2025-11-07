<template>
  <div 
    :class="['user-avatar-wrapper', sizeClass]" 
    :style="avatarStyle"
  >
    <img 
      v-if="avatarUrl && !imageError"
      :src="avatarUrl"
      :alt="username || '用户头像'"
      @error="handleImageError"
    />
    <span v-else class="avatar-text">{{ avatarText }}</span>
  </div>
</template>

<script setup>
import { computed, ref } from 'vue'

const props = defineProps({
  avatar: {
    type: String,
    default: ''
  },
  username: {
    type: String,
    default: ''
  },
  nickname: {
    type: String,
    default: ''
  },
  size: {
    type: String,
    default: 'medium' // small, medium, large
  }
})

const imageError = ref(false)

const sizeClass = computed(() => `avatar-${props.size}`)

const avatarUrl = computed(() => {
  if (!props.avatar) return ''
  // 如果是相对路径，添加完整 URL
  if (props.avatar.startsWith('/uploads')) {
    return `http://localhost:8080${props.avatar}`
  }
  // 如果是完整 URL（OSS），直接返回
  if (props.avatar.startsWith('http')) {
    return props.avatar
  }
  return props.avatar
})

const avatarText = computed(() => {
  const name = props.nickname || props.username || ''
  if (name) {
    return name.charAt(0).toUpperCase()
  }
  return 'U'
})

const avatarStyle = computed(() => {
  if (!avatarUrl.value || imageError.value) {
    return {
      background: 'linear-gradient(135deg, #ffb6c1 0%, #ffa07a 50%, #ffb347 100%)',
      color: '#fff'
    }
  }
  return {}
})

const handleImageError = () => {
  imageError.value = true
}
</script>

<style scoped>
.user-avatar-wrapper {
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(255, 182, 193, 0.4);
}

.user-avatar-wrapper img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.avatar-text {
  font-weight: 600;
  user-select: none;
}

.avatar-small {
  width: 32px;
  height: 32px;
  font-size: 12px;
}

.avatar-medium {
  width: 40px;
  height: 40px;
  font-size: 16px;
}

.avatar-large {
  width: 48px;
  height: 48px;
  font-size: 18px;
}

.user-avatar-wrapper:hover {
  transform: scale(1.05);
  transition: transform 0.2s;
}
</style>


