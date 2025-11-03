import request from './request'

export const uploadAPI = {
  // 上传图片
  uploadImage(file, onProgress) {
    const formData = new FormData()
    formData.append('image', file)
    
    return request({
      url: '/upload/image',
      method: 'post',
      data: formData,
      headers: {
        'Content-Type': 'multipart/form-data'
      },
      onUploadProgress: (progressEvent) => {
        if (onProgress && progressEvent.total) {
          const percentCompleted = Math.round((progressEvent.loaded * 100) / progressEvent.total)
          onProgress(percentCompleted)
        }
      }
    })
  },

  // 删除图片
  deleteImage(url) {
    return request({
      url: '/upload/image',
      method: 'delete',
      data: { url }
    })
  }
}



