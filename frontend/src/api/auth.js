import request from './request'

export const authAPI = {
  // 注册
  register(data) {
    return request.post('/auth/register', data)
  },

  // 登录
  login(data) {
    return request.post('/auth/login', data)
  },

  // 获取个人信息
  getProfile() {
    return request.get('/profile')
  }
}





