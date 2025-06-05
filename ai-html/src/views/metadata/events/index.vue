<template>
  <div class="events-container">
    <div class="table-header">
      <a-space>
        <a-input-search
          v-model:value="searchKeyword"
          placeholder="搜索事件名称或描述"
          style="width: 300px"
          @search="onSearch"
          @change="onSearchChange"
        />
        <a-button type="primary" @click="showCreateModal">
          <template #icon><PlusOutlined /></template>
          添加事件
        </a-button>
      </a-space>
    </div>

    <a-table
      :columns="columns"
      :data-source="eventList"
      :loading="loading"
      row-key="id"
      :pagination="pagination"
      @change="handleTableChange"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <a-space>
            <a-button type="link" @click="showEditModal(record)">编辑</a-button>
            <a-button type="link" @click="showAttributeModal(record)">查看事件属性</a-button>
            <a-popconfirm
              title="确定要删除这个事件吗？"
              @confirm="handleDelete(record.id)"
            >
              <a-button type="link" danger>删除</a-button>
            </a-popconfirm>
          </a-space>
        </template>
      </template>
    </a-table>

    <!-- 创建/编辑事件弹窗 -->
    <a-modal
      :title="modalTitle"
      v-model:open="modalVisible"
      @ok="handleModalOk"
      @cancel="handleModalCancel"
      :confirmLoading="modalLoading"
    >
      <a-form
        ref="formRef"
        :model="formState"
        :rules="rules"
        :label-col="{ span: 6 }"
        :wrapper-col="{ span: 16 }"
      >
        <a-form-item label="事件名称" name="name">
          <a-input v-model:value="formState.name" />
        </a-form-item>
        <a-form-item label="描述" name="description">
          <a-textarea v-model:value="formState.description" :rows="4" />
        </a-form-item>
      </a-form>
    </a-modal>

    <!-- 事件属性弹窗 -->
    <a-modal
      title="事件属性"
      v-model:open="attributeModalVisible"
      @ok="handleAttributeModalOk"
      @cancel="handleAttributeModalCancel"
      :confirmLoading="attributeModalLoading"
      width="800px"
    >
      <a-table
        :columns="attributeColumns"
        :data-source="selectedEventAttributes"
        :loading="attributeLoading"
        row-key="id"
        :pagination="false"
      >
      </a-table>
    </a-modal>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted } from 'vue'
import { PlusOutlined } from '@ant-design/icons-vue'
import { message } from 'ant-design-vue'
import type { FormInstance } from 'ant-design-vue'
import request from '@/utils/request'

interface EventInfo {
  id: number
  event_name: string
  event_desc: string
  is_deleted: boolean
  created_at: string
  updated_at: string
}

interface AttributeInfo {
  id: number
  attr_name: string
  attr_type: string
  attr_desc: string
  is_deleted: boolean
  created_at: string
  updated_at: string
}

interface ApiResponse<T> {
  code: number
  message: string
  data: T
}

interface EventListResponse {
  events: EventInfo[]
  pagination: {
    page_num: number
    page_size: number
    total: number
    total_page: number
  }
}

interface EventDetailResponse {
  id: number
  event_name: string
  event_desc: string
  is_deleted: boolean
  created_at: string
  updated_at: string
  attributes: AttributeInfo[]
}

const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    key: 'id',
  },
  {
    title: '事件名称',
    dataIndex: 'event_name',
    key: 'event_name',
  },
  {
    title: '描述',
    dataIndex: 'event_desc',
    key: 'event_desc',
    ellipsis: true,
  },
  {
    title: '创建时间',
    dataIndex: 'created_at',
    key: 'created_at',
  },
  {
    title: '操作',
    key: 'action',
    width: 250,
  },
]

const attributeColumns = [
  {
    title: '属性名称',
    dataIndex: 'attr_name',
    key: 'attr_name',
  },
  {
    title: '属性类型',
    dataIndex: 'attr_type',
    key: 'attr_type',
  },
  {
    title: '描述',
    dataIndex: 'attr_desc',
    key: 'attr_desc',
    ellipsis: true,
  },
  {
    title: '创建时间',
    dataIndex: 'created_at',
    key: 'created_at',
  },
]

const eventList = ref<EventInfo[]>([])
const loading = ref(false)
const modalVisible = ref(false)
const modalLoading = ref(false)
const modalTitle = ref('添加事件')
const editMode = ref(false)
const formRef = ref<FormInstance>()

const formState = reactive({
  id: undefined as number | undefined,
  name: '',
  description: '',
})

const rules = {
  name: [{ required: true, message: '请输入事件名称!' }],
  description: [{ required: true, message: '请输入描述!' }],
}

// 属性配置相关
const attributeModalVisible = ref(false)
const attributeModalLoading = ref(false)
const attributeLoading = ref(false)
const selectedEventAttributes = ref<AttributeInfo[]>([])
const selectedEventId = ref<number>()

// 分页配置
const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0,
  showTotal: (total: number) => `共 ${total} 条`,
  showSizeChanger: true,
  showQuickJumper: true,
})

// 搜索关键词
const searchKeyword = ref('')

// 获取事件列表
const fetchEventList = async (params = { page_num: 1, page_size: 10, keyword: '' }) => {
  loading.value = true
  try {
    const response = await request.get<any, ApiResponse<EventListResponse>>('/events', {
      params: {
        page_num: params.page_num,
        page_size: params.page_size,
        keyword: params.keyword,
      }
    })
    console.log('Events response:', response)
    if (response.code === 200 && response.data?.events) {
      eventList.value = response.data.events
      // 更新分页信息
      pagination.current = response.data.pagination.page_num
      pagination.pageSize = response.data.pagination.page_size
      pagination.total = response.data.pagination.total
      console.log('Event list:', eventList.value)
    } else {
      message.error(response.message || '获取事件列表失败')
    }
  } catch (error: any) {
    console.error('Error fetching events:', error)
    message.error(error.message || '获取事件列表失败')
  } finally {
    loading.value = false
  }
}

// 显示创建事件弹窗
const showCreateModal = () => {
  editMode.value = false
  modalTitle.value = '添加事件'
  formState.id = undefined
  formState.name = ''
  formState.description = ''
  modalVisible.value = true
}

// 显示编辑事件弹窗
const showEditModal = (record: EventInfo) => {
  editMode.value = true
  modalTitle.value = '编辑事件'
  formState.id = record.id
  formState.name = record.event_name
  formState.description = record.event_desc
  modalVisible.value = true
}

// 处理弹窗确认
const handleModalOk = async () => {
  try {
    await formRef.value?.validate()
    modalLoading.value = true

    const requestData = {
      event_name: formState.name,
      event_desc: formState.description,
    }

    let response
    if (editMode.value) {
      response = await request.post<any, ApiResponse<any>>(`/events/${formState.id}`, requestData)
    } else {
      response = await request.post<any, ApiResponse<any>>('/events', requestData)
    }

    if (response.code === 200) {
      message.success(editMode.value ? '编辑事件成功' : '创建事件成功')
      modalVisible.value = false
      fetchEventList()
    } else {
      message.error(response.message || (editMode.value ? '编辑事件失败' : '创建事件失败'))
    }
  } catch (error: any) {
    if (error.message) {
      message.error(error.message)
    }
  } finally {
    modalLoading.value = false
  }
}

// 处理弹窗取消
const handleModalCancel = () => {
  modalVisible.value = false
  formRef.value?.resetFields()
}

// 处理删除事件
const handleDelete = async (id: number) => {
  try {
    const response = await request.post<any, ApiResponse<any>>(`/events/${id}/delete`)
    if (response.code === 200) {
      message.success('删除事件成功')
      fetchEventList()
    } else {
      message.error(response.message || '删除事件失败')
    }
  } catch (error: any) {
    message.error(error.message || '删除事件失败')
  }
}

// 显示事件属性弹窗
const showAttributeModal = async (record: EventInfo) => {
  selectedEventId.value = record.id
  attributeModalVisible.value = true
  attributeLoading.value = true
  try {
    const response = await request.get<any, ApiResponse<EventDetailResponse>>(`/events/${record.id}`, {
      params: {
        page_num: '',  // 不传分页参数，获取所有数据
        page_size: '',
      }
    })
    if (response.code === 200 && response.data) {
      selectedEventAttributes.value = response.data.attributes || []
    } else {
      message.error(response.message || '获取事件属性失败')
    }
  } catch (error: any) {
    message.error(error.message || '获取事件属性失败')
  } finally {
    attributeLoading.value = false
  }
}

// 处理事件属性弹窗确认
const handleAttributeModalOk = async () => {
  try {
    attributeModalLoading.value = true
    const response = await request.post<any, ApiResponse<any>>(`/events/${selectedEventId.value}/attributes`, {
      attributes: selectedEventAttributes.value
    })
    if (response.code === 200) {
      message.success('保存事件属性成功')
      attributeModalVisible.value = false
    } else {
      message.error(response.message || '保存事件属性失败')
    }
  } catch (error: any) {
    message.error(error.message || '保存事件属性失败')
  } finally {
    attributeModalLoading.value = false
  }
}

// 处理事件属性弹窗取消
const handleAttributeModalCancel = () => {
  attributeModalVisible.value = false
  selectedEventAttributes.value = []
  selectedEventId.value = undefined
}

// 处理搜索
const onSearch = () => {
  fetchEventList({
    page_num: 1,
    page_size: pagination.pageSize,
    keyword: searchKeyword.value,
  })
}

// 处理搜索框值变化
const onSearchChange = () => {
  if (!searchKeyword.value) {
    fetchEventList({
      page_num: 1,
      page_size: pagination.pageSize,
      keyword: '',
    })
  }
}

// 处理表格变化（分页、排序等）
const handleTableChange = (pag: any) => {
  fetchEventList({
    page_num: pag.current,
    page_size: pag.pageSize,
    keyword: searchKeyword.value,
  })
}

onMounted(() => {
  fetchEventList()
})
</script>

<style scoped>
.events-container {
  padding: 24px;
}

.table-header {
  margin-bottom: 16px;
  display: flex;
  justify-content: space-between;
}
</style> 