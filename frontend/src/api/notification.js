import request from './request'

export const notificationAPI = {
  // 获取通知列表
  getNotifications(params) {
    return request.get('/notifications', { params })
  },

  // 获取未读通知数量
  getUnreadCount() {
    return request.get('/notifications/unread-count')
  },

  // 标记通知为已读
  markAsRead(id) {
    return request.patch(`/notifications/${id}/read`)
  },

  // 标记所有通知为已读
  markAllAsRead() {
    return request.patch('/notifications/read-all')
  },

  // 删除通知
  deleteNotification(id) {
    return request.delete(`/notifications/${id}`)
  }
}

