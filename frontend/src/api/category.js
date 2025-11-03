import request from './request'

export const categoryAPI = {
  // 获取所有分类
  getCategories() {
    return request.get('/categories')
  },

  // 获取单个分类
  getCategory(id) {
    return request.get(`/categories/${id}`)
  },

  // 创建分类
  createCategory(data) {
    return request.post('/categories', data)
  },

  // 更新分类
  updateCategory(id, data) {
    return request.put(`/categories/${id}`, data)
  },

  // 删除分类
  deleteCategory(id) {
    return request.delete(`/categories/${id}`)
  }
}





