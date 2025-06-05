<template>
  <div class="user-container">
    <div class="table-header">
      <a-button type="primary" @click="showCreateModal">
        <template #icon><PlusOutlined /></template>
        添加用户
      </a-button>
    </div>

    <a-table
      :columns="columns"
      :data-source="userList"
      :loading="loading"
      row-key="id"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <a-space>
            <a-button type="link" @click="showEditModal(record)">编辑</a-button>
            <a-popconfirm
              title="确定要删除这个用户吗？"
              @confirm="handleDelete(record.id)"
            >
              <a-button type="link" danger>删除</a-button>
            </a-popconfirm>
          </a-space>
        </template>
      </template>
    </a-table>

    <!-- 创建/编辑用户弹窗 -->
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
        <a-form-item label="用户名" name="username">
          <a-input v-model:value="formState.username" />
        </a-form-item>
        <a-form-item 
          label="密码" 
          name="password"
          :rules="[{ required: !editMode, message: '请输入密码!' }]"
        >
          <a-input-password v-model:value="formState.password" />
        </a-form-item>
        <a-form-item label="角色" name="role">
          <a-select v-model:value="formState.role">
            <a-select-option value="admin">管理员</a-select-option>
            <a-select-option value="user">普通用户</a-select-option>
          </a-select>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted, reactive } from 'vue'
import { PlusOutlined } from '@ant-design/icons-vue'
import { message } from 'ant-design-vue'
import type { FormInstance } from 'ant-design-vue'
import request from '../../utils/request'

interface UserInfo {
  id: number
  username: string
  role: string
  created_at: string
  updated_at: string
}

interface FormState {
  id?: number
  username: string
  password: string
  role: string
}

const columns = [
  {
    title: 'ID',
    dataIndex: 'id',
    key: 'id',
  },
  {
    title: '用户名',
    dataIndex: 'username',
    key: 'username',
  },
  {
    title: '角色',
    dataIndex: 'role',
    key: 'role',
  },
  {
    title: '创建时间',
    dataIndex: 'created_at',
    key: 'created_at',
  },
  {
    title: '更新时间',
    dataIndex: 'updated_at',
    key: 'updated_at',
  },
  {
    title: '操作',
    key: 'action',
  },
]

const userList = ref<UserInfo[]>([])
const loading = ref(false)
const modalVisible = ref(false)
const modalLoading = ref(false)
const modalTitle = ref('添加用户')
const editMode = ref(false)
const formRef = ref<FormInstance>()

const formState = reactive<FormState>({
  username: '',
  password: '',
  role: 'user',
})

const rules = {
  username: [{ required: true, message: '请输入用户名!' }],
  role: [{ required: true, message: '请选择角色!' }],
}

// 获取用户列表
const fetchUserList = async () => {
  loading.value = true
  try {
    const response = await request.get('/users')
    userList.value = response.data
  } catch (error: any) {
    message.error(error.message || '获取用户列表失败')
  } finally {
    loading.value = false
  }
}

// 显示创建用户弹窗
const showCreateModal = () => {
  editMode.value = false
  modalTitle.value = '添加用户'
  formRef.value?.resetFields()
  formState.role = 'user'
  formState.id = undefined
  formState.username = ''
  formState.password = ''
  modalVisible.value = true
}

// 显示编辑用户弹窗
const showEditModal = (record: UserInfo) => {
  editMode.value = true
  modalTitle.value = '编辑用户'
  formRef.value?.resetFields()
  formState.id = record.id
  formState.username = record.username
  formState.password = ''
  formState.role = record.role
  modalVisible.value = true
}

// 处理弹窗确认
const handleModalOk = async () => {
  try {
    await formRef.value?.validate()
    modalLoading.value = true

    if (editMode.value) {
      // 编辑用户
      await request.post(`/users/${formState.id}`, {
        username: formState.username,
        password: formState.password || undefined,
        role: formState.role,
      })
      message.success('编辑用户成功')
    } else {
      // 创建用户
      await request.post('/users', formState)
      message.success('创建用户成功')
    }

    modalVisible.value = false
    fetchUserList()
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
  formState.id = undefined
  formState.username = ''
  formState.password = ''
  formState.role = 'user'
}

// 处理删除用户
const handleDelete = async (id: number) => {
  try {
    await request.post(`/users/${id}/delete`)
    message.success('删除用户成功')
    fetchUserList()
  } catch (error: any) {
    message.error(error.message || '删除用户失败')
  }
}

onMounted(() => {
  fetchUserList()
})
</script>

<style scoped>
.user-container {
  padding: 24px;
}

.table-header {
  margin-bottom: 16px;
}
</style> 