<template>
  <a-layout class="layout">
    <a-layout-sider
      v-model:collapsed="collapsed"
      :trigger="null"
      collapsible
      class="sider"
    >
      <div class="logo">
        <h1 v-if="!collapsed">元数据管理系统</h1>
        <h1 v-else>AI</h1>
      </div>
      <a-menu
        v-model:selectedKeys="selectedKeys"
        theme="dark"
        mode="inline"
        :items="menuItems"
        @click="handleMenuClick"
      />
      <div class="trigger-wrapper">
        <a-button
          type="text"
          class="trigger-button"
          @click="() => (collapsed = !collapsed)"
        >
          <LeftOutlined v-if="!collapsed" />
          <RightOutlined v-else />
        </a-button>
      </div>
    </a-layout-sider>
    <a-layout>
      <a-layout-header class="header">
        <a-space>
          <a-dropdown>
            <a class="user-dropdown" @click.prevent>
              <a-avatar>
                <template #icon><UserOutlined /></template>
              </a-avatar>
              <span class="username">{{ username }}</span>
            </a>
            <template #overlay>
              <a-menu>
                <a-menu-item key="logout" @click="handleLogout">
                  <LogoutOutlined />
                  <span>退出登录</span>
                </a-menu-item>
              </a-menu>
            </template>
          </a-dropdown>
        </a-space>
      </a-layout-header>
      <a-layout-content class="content">
        <router-view></router-view>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script lang="ts" setup>
import { ref, h, onMounted, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import {
  UserOutlined,
  LogoutOutlined,
  DashboardOutlined,
  DatabaseOutlined,
  TeamOutlined,
  LeftOutlined,
  RightOutlined
} from '@ant-design/icons-vue'
import { message } from 'ant-design-vue'

const router = useRouter()
const route = useRoute()
const collapsed = ref<boolean>(false)
const selectedKeys = ref<string[]>(['dashboard'])
const username = ref<string>('')

// 获取用户信息
const getUserInfo = () => {
  const userInfoStr = localStorage.getItem('userInfo')
  if (userInfoStr) {
    const userInfo = JSON.parse(userInfoStr)
    username.value = userInfo.username
  }
}

// 根据当前路由设置选中的菜单项
const setSelectedKey = () => {
  const path = route.path
  const mainPath = path.split('/')[1] || 'dashboard'
  selectedKeys.value = [mainPath]
}

const menuItems = [
  {
    key: 'dashboard',
    icon: () => h(DashboardOutlined),
    label: '仪表盘',
  },
  {
    key: 'metadata',
    icon: () => h(DatabaseOutlined),
    label: '元数据管理',
  },
  {
    key: 'user',
    icon: () => h(TeamOutlined),
    label: '用户管理',
  },
]

const handleMenuClick = (menu: any) => {
  if (menu.key === 'metadata') {
    router.push('/metadata/events')
  } else {
    router.push(`/${menu.key}`)
  }
}

const handleLogout = () => {
  localStorage.removeItem('token')
  message.success('退出成功')
  router.push('/login')
}

onMounted(() => {
  getUserInfo()
  setSelectedKey()
})

// 监听路由变化
watch(
  () => route.path,
  () => {
    setSelectedKey()
  }
)
</script>

<style scoped>
.layout {
  min-height: 100vh;
}

.logo {
  height: 64px;
  line-height: 64px;
  text-align: center;
  color: white;
  font-size: 16px;
  overflow: hidden;
}

.logo h1 {
  color: white;
  font-size: 16px;
  margin: 0;
}

.header {
  background: #fff;
  padding: 0 24px;
  display: flex;
  justify-content: flex-end;
  align-items: center;
}

.content {
  margin: 24px 16px;
  padding: 24px;
  background: #fff;
  min-height: 280px;
}

.user-dropdown {
  display: flex;
  align-items: center;
  color: rgba(0, 0, 0, 0.85);
}

.username {
  margin-left: 8px;
}

.sider {
  position: relative;
}

.trigger-wrapper {
  position: absolute;
  right: -16px;
  top: 50%;
  transform: translateY(-50%);
  z-index: 1;
}

.trigger-button {
  width: 16px;
  height: 16px;
  color: #fff;
  font-size: 12px;
  text-align: center;
  background-color: #001529;
  border-radius: 50%;
  border: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s;
  padding: 0;
  min-width: 16px;
}

.trigger-button:hover {
  color: #1890ff;
  background-color: #002140;
}

.trigger-button :deep(.anticon) {
  font-size: 10px;
  transition: transform 0.3s;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
}

.trigger-button:hover :deep(.anticon) {
  transform: scale(1.2);
}
</style> 