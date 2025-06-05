import axios from 'axios'
import { message } from 'ant-design-vue'
import router from '../router'

// 创建axios实例
const request = axios.create({
  baseURL: 'http://localhost:8080/api', // 添加/api前缀
  timeout: 15000,
})

// 请求拦截器
request.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
request.interceptors.response.use(
  (response) => {
    // 如果是登录接口，保存token
    if (response.config.url === '/login') {
      const { data } = response.data
      if (data && data.token) {
        localStorage.setItem('token', data.token)
      }
    }
    // 如果是登出接口，清除token和用户信息
    else if (response.config.url === '/logout') {
      localStorage.removeItem('token')
      localStorage.removeItem('userInfo')
      router.push('/login')
    }
    return response.data
  },
  (error) => {
    if (error.response) {
      switch (error.response.status) {
        case 401:
          // token过期或未登录
          localStorage.removeItem('token')
          localStorage.removeItem('userInfo')
          message.error('登录已过期，请重新登录')
          router.push('/login')
          break
        case 403:
          message.error('没有权限访问')
          break
        case 404:
          message.error('请求的资源不存在')
          break
        case 500:
          message.error('服务器错误')
          break
        default:
          message.error(error.response.data?.message || '请求失败')
      }
    } else {
      message.error('网络错误，请检查网络连接')
    }
    return Promise.reject(error)
  }
)

export default request 