import request from './request'

export const articleAPI = {
  // 获取文章列表
  getArticles(params) {
    return request.get('/articles', { params })
  },

  // 获取单篇文章
  getArticle(id) {
    return request.get(`/articles/${id}`)
  },

  // 创建文章
  createArticle(data) {
    return request.post('/articles', data)
  },

  // 更新文章
  updateArticle(id, data) {
    return request.put(`/articles/${id}`, data)
  },

  // 删除文章
  deleteArticle(id) {
    return request.delete(`/articles/${id}`)
  },

  // 获取我的文章
  getMyArticles(params) {
    return request.get('/my-articles', { params })
  }
}





