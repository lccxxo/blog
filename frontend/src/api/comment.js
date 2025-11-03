import request from './request'

export const commentAPI = {
  // 获取文章的评论列表
  getArticleComments(articleId) {
    return request.get(`/articles/${articleId}/comments`)
  },

  // 获取单个评论
  getComment(id) {
    return request.get(`/comments/${id}`)
  },

  // 创建评论
  createComment(data) {
    return request.post('/comments', data)
  },

  // 更新评论
  updateComment(id, data) {
    return request.put(`/comments/${id}`, data)
  },

  // 删除评论
  deleteComment(id) {
    return request.delete(`/comments/${id}`)
  },

  // 获取我的评论
  getMyComments(params) {
    return request.get('/my-comments', { params })
  }
}





