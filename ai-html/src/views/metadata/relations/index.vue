<template>
  <div class="relations-container">
    <div class="table-header">
      <a-button type="primary" @click="showCreateModal">
        <template #icon><PlusOutlined /></template>
        添加关联
      </a-button>
    </div>

    <a-table
      :columns="columns"
      :data-source="relationList"
      :loading="loading"
      :pagination="pagination"
      @change="handleTableChange"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <a-space>
            <a-popconfirm
              title="确定要删除这个关联关系吗？"
              ok-text="确定"
              cancel-text="取消"
              @confirm="() => handleDelete(record.id)"
            >
              <a-button type="link" danger>删除</a-button>
            </a-popconfirm>
          </a-space>
        </template>
      </template>
    </a-table>

    <!-- 创建关联弹窗 -->
    <a-modal
      v-model:open="modalVisible"
      title="添加关联"
      :confirm-loading="modalLoading"
      @ok="handleModalOk"
      @cancel="handleModalCancel"
    >
      <a-form
        ref="formRef"
        :model="formState"
        :rules="rules"
        :label-col="{ span: 6 }"
        :wrapper-col="{ span: 16 }"
      >
        <a-form-item label="事件" name="event_id">
          <a-select
            v-model:value="formState.event_id"
            placeholder="请选择事件"
            :options="eventOptions"
            :field-names="{ label: 'event_name', value: 'id' }"
          />
        </a-form-item>
        <a-form-item label="属性" name="attr_id">
          <a-select
            v-model:value="formState.attr_id"
            placeholder="请选择属性"
            :options="attrOptions"
            :field-names="{ label: 'attr_name', value: 'id' }"
          />
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive, onMounted } from 'vue'
import { PlusOutlined } from '@ant-design/icons-vue'
import { message } from 'ant-design-vue'
import type { FormInstance } from 'ant-design-vue'
import request from '@/utils/request'

interface RelationInfo {
  id: number
  event_id: number
  attr_id: number
  event_name: string
  attr_name: string
  created_at: string
}

interface EventOption {
  id: number
  event_name: string
}

interface AttrOption {
  id: number
  attr_name: string
}

interface ApiResponse<T> {
  code: number
  message: string
  data: T
}

interface RelationListResponse {
  relations: RelationInfo[]
  pagination: {
    page_num: number
    page_size: number
    total: number
    total_page: number
  }
}

interface EventListResponse {
  events: EventOption[]
}

interface AttributeListResponse {
  attributes: AttrOption[]
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
    title: '属性名称',
    dataIndex: 'attr_name',
    key: 'attr_name',
  },
  {
    title: '创建时间',
    dataIndex: 'created_at',
    key: 'created_at',
  },
  {
    title: '操作',
    key: 'action',
    width: 200,
  },
]

const relationList = ref<RelationInfo[]>([])
const loading = ref(false)
const modalVisible = ref(false)
const modalLoading = ref(false)
const formRef = ref<FormInstance>()
const eventOptions = ref<EventOption[]>([])
const attrOptions = ref<AttrOption[]>([])

const formState = reactive({
  event_id: undefined as number | undefined,
  attr_id: undefined as number | undefined,
})

const rules = {
  event_id: [{ required: true, message: '请选择事件!' }],
  attr_id: [{ required: true, message: '请选择属性!' }],
}

// 分页配置
const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0,
  showTotal: (total: number) => `共 ${total} 条`,
  showSizeChanger: true,
  showQuickJumper: true,
})

// 获取关联列表
const fetchRelationList = async (params = { page_num: 1, page_size: 10 }) => {
  loading.value = true
  try {
    const response = await request.get<any, ApiResponse<RelationListResponse>>('/relations', {
      params: {
        page_num: params.page_num,
        page_size: params.page_size,
      }
    })
    if (response.code === 200 && response.data?.relations) {
      relationList.value = response.data.relations
      // 更新分页信息
      pagination.current = response.data.pagination.page_num
      pagination.pageSize = response.data.pagination.page_size
      pagination.total = response.data.pagination.total
    } else {
      message.error(response.message || '获取关联列表失败')
    }
  } catch (error: any) {
    console.error('Error fetching relations:', error)
    message.error(error.message || '获取关联列表失败')
  } finally {
    loading.value = false
  }
}

// 获取事件列表（用于下拉选择）
const fetchEventOptions = async () => {
  try {
    const response = await request.get<any, ApiResponse<EventListResponse>>('/events')
    if (response.code === 200 && response.data?.events) {
      eventOptions.value = response.data.events
    }
  } catch (error: any) {
    console.error('Error fetching events:', error)
    message.error('获取事件列表失败')
  }
}

// 获取属性列表（用于下拉选择）
const fetchAttrOptions = async () => {
  try {
    const response = await request.get<any, ApiResponse<AttributeListResponse>>('/attributes')
    if (response.code === 200 && response.data?.attributes) {
      attrOptions.value = response.data.attributes
    }
  } catch (error: any) {
    console.error('Error fetching attributes:', error)
    message.error('获取属性列表失败')
  }
}

// 显示创建关联弹窗
const showCreateModal = () => {
  formState.event_id = undefined
  formState.attr_id = undefined
  modalVisible.value = true
  // 获取事件和属性列表
  fetchEventOptions()
  fetchAttrOptions()
}

// 处理弹窗确认
const handleModalOk = async () => {
  try {
    await formRef.value?.validate()
    modalLoading.value = true

    const response = await request.post<any, ApiResponse<any>>('/relations', {
      event_id: formState.event_id,
      attr_id: formState.attr_id,
    })

    if (response.code === 200) {
      message.success('创建关联成功')
      modalVisible.value = false
      fetchRelationList()
    } else {
      message.error(response.message || '创建关联失败')
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

// 处理删除关联
const handleDelete = async (id: number) => {
  try {
    const response = await request.post<any, ApiResponse<any>>(`/relations/${id}/delete`)
    if (response.code === 200) {
      message.success('删除关联成功')
      fetchRelationList()
    } else {
      message.error(response.message || '删除关联失败')
    }
  } catch (error: any) {
    message.error(error.message || '删除关联失败')
  }
}

// 处理表格变化（分页、排序等）
const handleTableChange = (pag: any) => {
  fetchRelationList({
    page_num: pag.current,
    page_size: pag.pageSize,
  })
}

onMounted(() => {
  fetchRelationList()
})
</script>

<style scoped>
.relations-container {
  padding: 24px;
}

.table-header {
  margin-bottom: 16px;
}
</style> 