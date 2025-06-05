<template>
  <div class="metadata-container">
    <a-menu
      v-model:selectedKeys="selectedKeys"
      mode="inline"
      class="metadata-menu"
      @click="handleMenuClick"
    >
      <a-menu-item key="events">
        <template #icon>
          <ThunderboltOutlined />
        </template>
        事件管理
      </a-menu-item>
      <a-menu-item key="attributes">
        <template #icon>
          <TagsOutlined />
        </template>
        属性管理
      </a-menu-item>
      <a-menu-item key="relations">
        <template #icon>
          <NodeIndexOutlined />
        </template>
        关联管理
      </a-menu-item>
    </a-menu>
    <div class="metadata-content">
      <router-view></router-view>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { ThunderboltOutlined, TagsOutlined, NodeIndexOutlined } from '@ant-design/icons-vue'

const router = useRouter()
const route = useRoute()
const selectedKeys = ref<string[]>(['events'])

// 根据当前路由设置选中的菜单项
const setSelectedKey = () => {
  const path = route.path
  const key = path.split('/').pop() || 'events'
  selectedKeys.value = [key]
}

// 处理菜单点击事件
const handleMenuClick = ({ key }: { key: string }) => {
  router.push(`/metadata/${key}`)
}

// 监听路由变化
watch(
  () => route.path,
  () => {
    setSelectedKey()
  }
)

onMounted(() => {
  // 如果当前路径是 /metadata，重定向到 /metadata/events
  if (route.path === '/metadata') {
    router.push('/metadata/events')
  }
  setSelectedKey()
})
</script>

<style scoped>
.metadata-container {
  display: flex;
  height: 100%;
}

.metadata-menu {
  width: 200px;
  border-right: 1px solid #f0f0f0;
}

.metadata-content {
  flex: 1;
  padding: 0;
  background: #fff;
}
</style> 