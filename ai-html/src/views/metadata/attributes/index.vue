<template>
  <div class="attributes-container">
    <div class="table-header">
      <a-space>
        <a-input-search
          v-model:value="searchKeyword"
          placeholder="搜索属性名称或描述"
          style="width: 300px"
          @search="onSearch"
          @change="onSearchChange"
        />
        <a-button type="primary" @click="showCreateModal">
          <template #icon><PlusOutlined /></template>
          添加属性
        </a-button>
      </a-space>
    </div>

    <a-table
      :columns="columns"
      :data-source="attributeList"
      :loading="loading"
      row-key="id"
      :pagination="pagination"
      @change="handleTableChange"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <a-space>
            <a-button type="link" @click="showEditModal(record)">编辑</a-button>
            <a-popconfirm
              title="确定要删除这个属性吗？"
              @confirm="handleDelete(record.id)"
            >
              <a-button type="link" danger>删除</a-button>
            </a-popconfirm>
          </a-space>
        </template>
      </template>
    </a-table>

    <!-- 创建/编辑属性弹窗 -->
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
        <a-form-item label="属性名称" name="name">
          <a-input v-model:value="formState.name" />
        </a-form-item>
        <a-form-item label="属性类型" name="type">
          <a-select v-model:value="formState.type" placeholder="请选择属性类型">
            <a-select-option
              v-for="type in attributeTypes"
              :key="type.value"
              :value="type.value"
            >
              {{ type.label }}
            </a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item label="描述" name="description">
          <a-textarea v-model:value="formState.description" :rows="4" />
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

interface AttributeInfo {
  id: number
  attr_name: string
  attr_type: string
  attr_desc: string
  is_deleted: boolean
  created_at: string
  updated_at: string
}

interface AttributeType {
  value: string
  label: string
}

interface ApiResponse<T> {
  code: number
  message: string
  data: T
}

interface AttributeListResponse {
  attributes: AttributeInfo[]
  pagination: {
    page_num: number
    page_size: number
    total: number
    total_page: number
  }
}

const attributeList = ref<AttributeInfo[]>([])
const loading = ref(false)
const modalVisible = ref(false)
const modalLoading = ref(false)
const modalTitle = ref('添加属性')
const editMode = ref(false)
const formRef = ref<FormInstance>()
const attributeTypes = ref<AttributeType[]>([])

const formState = reactive({
  id: undefined as number | undefined,
  name: '',
  type: '',
  description: '',
})

const rules = {
  name: [{ required: true, message: '请输入属性名称!' }],
  type: [{ required: true, message: '请选择属性类型!' }],
  description: [{ required: true, message: '请输入描述!' }],
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

// 搜索关键词
const searchKeyword = ref('')

// 获取属性类型列表
const fetchAttributeTypes = async () => {
  try {
    const response = await request.get<any, ApiResponse<{ types: AttributeType[] }>>('/attribute-types')
    if (response.code === 200) {
      attributeTypes.value = response.data.types
    }
  } catch (error: any) {
    message.error(error.message || '获取属性类型列表失败')
  }
}

// 获取属性类型的中文标签
const getTypeLabel = (type: string) => {
  const typeItem = attributeTypes.value.find(t => t.value === type)
  return typeItem ? typeItem.label : type
}

const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    key: 'id',
  },
  {
    title: '属性名称',
    dataIndex: 'attr_name',
    key: 'attr_name',
  },
  {
    title: '属性类型',
    dataIndex: 'attr_type',
    key: 'attr_type',
    customRender: ({ text }: { text: string }) => getTypeLabel(text),
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
  {
    title: '操作',
    key: 'action',
    width: 200,
  },
]

// 获取属性列表
const fetchAttributeList = async (params = { page_num: 1, page_size: 10, keyword: '' }) => {
  loading.value = true
  try {
    const response = await request.get<any, ApiResponse<AttributeListResponse>>('/attributes', {
      params: {
        page_num: params.page_num,
        page_size: params.page_size,
        keyword: params.keyword,
      }
    })
    console.log('Attributes response:', response)
    if (response.code === 200 && response.data?.attributes) {
      attributeList.value = response.data.attributes
      // 更新分页信息
      pagination.current = response.data.pagination.page_num
      pagination.pageSize = response.data.pagination.page_size
      pagination.total = response.data.pagination.total
      console.log('Attribute list:', attributeList.value)
    } else {
      message.error(response.message || '获取属性列表失败')
    }
  } catch (error: any) {
    console.error('Error fetching attributes:', error)
    message.error(error.message || '获取属性列表失败')
  } finally {
    loading.value = false
  }
}

// 显示创建属性弹窗
const showCreateModal = () => {
  editMode.value = false
  modalTitle.value = '添加属性'
  formState.id = undefined
  formState.name = ''
  formState.type = ''
  formState.description = ''
  modalVisible.value = true
  // 获取属性类型列表
  fetchAttributeTypes()
}

// 显示编辑属性弹窗
const showEditModal = (record: AttributeInfo) => {
  editMode.value = true
  modalTitle.value = '编辑属性'
  formState.id = record.id
  formState.name = record.attr_name
  formState.type = record.attr_type
  formState.description = record.attr_desc
  modalVisible.value = true
  // 获取属性类型列表
  fetchAttributeTypes()
}

// 处理弹窗确认
const handleModalOk = async () => {
  try {
    await formRef.value?.validate()
    modalLoading.value = true

    const requestData = {
      attr_name: formState.name,
      attr_type: formState.type,
      attr_desc: formState.description,
    }

    let response
    if (editMode.value) {
      response = await request.post<any, ApiResponse<any>>(`/attributes/${formState.id}`, requestData)
    } else {
      response = await request.post<any, ApiResponse<any>>('/attributes', requestData)
    }

    if (response.code === 200) {
      message.success(editMode.value ? '编辑属性成功' : '创建属性成功')
      modalVisible.value = false
      fetchAttributeList()
    } else {
      message.error(response.message || (editMode.value ? '编辑属性失败' : '创建属性失败'))
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

// 处理删除属性
const handleDelete = async (id: number) => {
  try {
    const response = await request.post<any, ApiResponse<any>>(`/attributes/${id}/delete`)
    if (response.code === 200) {
      message.success('删除属性成功')
      fetchAttributeList()
    } else {
      message.error(response.message || '删除属性失败')
    }
  } catch (error: any) {
    message.error(error.message || '删除属性失败')
  }
}

// 处理搜索
const onSearch = () => {
  fetchAttributeList({
    page_num: 1,
    page_size: pagination.pageSize,
    keyword: searchKeyword.value,
  })
}

// 处理搜索框值变化
const onSearchChange = () => {
  if (!searchKeyword.value) {
    fetchAttributeList({
      page_num: 1,
      page_size: pagination.pageSize,
      keyword: '',
    })
  }
}

// 处理表格变化（分页、排序等）
const handleTableChange = (pag: any) => {
  fetchAttributeList({
    page_num: pag.current,
    page_size: pag.pageSize,
    keyword: searchKeyword.value,
  })
}

onMounted(() => {
  fetchAttributeTypes()
  fetchAttributeList()
})
</script>

<style scoped>
.attributes-container {
  padding: 24px;
}

.table-header {
  margin-bottom: 16px;
  display: flex;
  justify-content: space-between;
}
</style> 