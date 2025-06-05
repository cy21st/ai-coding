<template>
  <div class="login-container">
    <a-card class="login-card" :bordered="false">
      <template #title>
        <div class="login-title">元数据管理后台</div>
      </template>
      <a-form
        :model="formState"
        name="login"
        @finish="onFinish"
        @finishFailed="onFinishFailed"
        autocomplete="off"
      >
        <a-form-item
          name="username"
          :rules="[{ required: true, message: '请输入用户名!' }]"
        >
          <a-input v-model:value="formState.username" placeholder="用户名">
            <template #prefix>
              <UserOutlined />
            </template>
          </a-input>
        </a-form-item>

        <a-form-item
          name="password"
          :rules="[{ required: true, message: '请输入密码!' }]"
        >
          <a-input-password v-model:value="formState.password" placeholder="密码">
            <template #prefix>
              <LockOutlined />
            </template>
          </a-input-password>
        </a-form-item>

        <a-form-item>
          <a-button type="primary" html-type="submit" :loading="loading" block>
            {{ loading ? '登录中...' : '登录' }}
          </a-button>
        </a-form-item>
      </a-form>
    </a-card>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive } from 'vue'
import { UserOutlined, LockOutlined } from '@ant-design/icons-vue'
import { message } from 'ant-design-vue'
import { useRouter } from 'vue-router'
import request from '../../utils/request'

interface LoginResponse {
  code: number
  message: string
  data: {
    token: string
    user: {
      id: number
      username: string
      role: string
    }
  }
}

const router = useRouter()
const loading = ref(false)

interface FormState {
  username: string
  password: string
}

const formState = reactive<FormState>({
  username: 'admin',
  password: '123456',
})

const onFinish = async (values: FormState) => {
  loading.value = true
  try {
    const response = await request.post<any, LoginResponse>('/login', values)
    if (response.code === 200) {
      // 存储token和用户信息
      localStorage.setItem('token', response.data.token)
      localStorage.setItem('userInfo', JSON.stringify(response.data.user))
      message.success('登录成功')
      router.push('/dashboard')
    } else {
      message.error(response.message || '登录失败')
    }
  } catch (error: any) {
    message.error(error.message || '登录失败，请检查用户名和密码')
  } finally {
    loading.value = false
  }
}

const onFinishFailed = (errorInfo: any) => {
  console.log('Failed:', errorInfo)
  message.error('请填写完整的登录信息')
}
</script>

<style scoped>
.login-container {
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #f0f2f5;
  background-image: url('data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSI1IiBoZWlnaHQ9IjUiPgo8cmVjdCB3aWR0aD0iNSIgaGVpZ2h0PSI1IiBmaWxsPSIjZmZmIj48L3JlY3Q+CjxyZWN0IHdpZHRoPSIxIiBoZWlnaHQ9IjEiIGZpbGw9IiNjY2MiPjwvcmVjdD4KPC9zdmc+');
}

.login-card {
  width: 100%;
  max-width: 400px;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
}

.login-title {
  font-size: 24px;
  font-weight: bold;
  color: rgba(0, 0, 0, 0.85);
  text-align: center;
  margin-bottom: 8px;
}

.login-card :deep(.ant-card-head) {
  border-bottom: none;
  padding: 32px 24px 0;
}

.login-card :deep(.ant-card-head-title) {
  padding: 0;
  text-align: center;
}

.login-card :deep(.ant-input-affix-wrapper) {
  height: 40px;
}

.login-card :deep(.ant-btn) {
  height: 40px;
}
</style> 