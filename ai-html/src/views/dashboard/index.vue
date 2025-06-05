<template>
  <div class="dashboard-container">
    <a-row :gutter="16">
      <a-col :span="6">
        <a-card class="stat-card" @click="router.push('/user')">
          <template #title>
            <span class="card-title">
              <TeamOutlined />
              用户总数
            </span>
          </template>
          <div class="card-value">{{ statistics.user_count }}</div>
        </a-card>
      </a-col>
      <a-col :span="6">
        <a-card class="stat-card" @click="router.push('/metadata/events')">
          <template #title>
            <span class="card-title">
              <AppstoreOutlined />
              事件总数
            </span>
          </template>
          <div class="card-value">{{ statistics.event_count }}</div>
        </a-card>
      </a-col>
      <a-col :span="6">
        <a-card class="stat-card" @click="router.push('/metadata/attributes')">
          <template #title>
            <span class="card-title">
              <DatabaseOutlined />
              属性总数
            </span>
          </template>
          <div class="card-value">{{ statistics.attr_count }}</div>
        </a-card>
      </a-col>
      <a-col :span="6">
        <a-card class="stat-card" @click="router.push('/metadata/relations')">
          <template #title>
            <span class="card-title">
              <LinkOutlined />
              关联总数
            </span>
          </template>
          <div class="card-value">{{ statistics.relation_count }}</div>
        </a-card>
      </a-col>
    </a-row>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import { message } from 'ant-design-vue'
import { useRouter } from 'vue-router'
import {
  TeamOutlined,
  AppstoreOutlined,
  DatabaseOutlined,
  LinkOutlined,
} from '@ant-design/icons-vue'
import request from '@/utils/request'

interface Statistics {
  user_count: number
  event_count: number
  attr_count: number
  relation_count: number
}

interface ApiResponse<T> {
  code: number
  message: string
  data: T
}

const router = useRouter()

const statistics = ref<Statistics>({
  user_count: 0,
  event_count: 0,
  attr_count: 0,
  relation_count: 0,
})

// 获取统计数据
const fetchStatistics = async () => {
  try {
    const response = await request.get<any, ApiResponse<Statistics>>('/statistics')
    if (response.code === 200) {
      statistics.value = response.data
    } else {
      message.error(response.message || '获取统计数据失败')
    }
  } catch (error: any) {
    console.error('Error fetching statistics:', error)
    message.error(error.message || '获取统计数据失败')
  }
}

onMounted(() => {
  fetchStatistics()
})
</script>

<style scoped>
.dashboard-container {
  padding: 24px;
}

.card-title {
  display: flex;
  align-items: center;
  font-size: 16px;
  color: rgba(0, 0, 0, 0.85);
}

.card-title :deep(.anticon) {
  margin-right: 8px;
  font-size: 20px;
}

.card-value {
  font-size: 36px;
  font-weight: bold;
  color: #1890ff;
  text-align: center;
  padding: 16px 0;
}

:deep(.ant-card) {
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.09);
}

:deep(.ant-card-head) {
  border-bottom: none;
  padding: 0 24px;
  min-height: 48px;
}

:deep(.ant-card-head-title) {
  padding: 12px 0;
}

:deep(.ant-card-body) {
  padding: 0 24px 24px;
}

.stat-card {
  cursor: pointer;
  transition: all 0.3s;
}

.stat-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}
</style> 