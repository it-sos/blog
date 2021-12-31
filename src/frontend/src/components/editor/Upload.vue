<template>
  <el-upload
      class="upload-demo"
      :action="action"
      :file-list="fileList"
      :on-success="handleSuccess">
    <el-button size="small" type="primary">上传</el-button>
    <template #tip>
      <div class="el-upload__tip">
        最大不能超过 5G ！
      </div>
    </template>
    <template #file="{ file }">
      <ul class="el-upload-list el-upload-list--text">
        <li class="el-upload-list__item is-success">
          <a class="el-upload-list__item-name" @click="insertDoc(file.file, file.name)"><i class="el-icon-document"></i>
            {{ file.name }} </a>
          <label class="el-upload-list__item-status-label">
            <i class="el-icon-upload-success el-icon-circle-check"></i>
          </label>
          <i class="el-icon-close" @click="removeFile(file.file)"></i>
          <i class="el-icon-close-tip">press delete to remove</i>
        </li>
      </ul>
    </template>
  </el-upload>
</template>
<script lang="ts">
import {defineComponent, inject, onMounted, ref} from "vue";
import utils from "@/common/utils";
import {ElMessage} from "element-plus";


interface FileList {
  name: string
  file: string
}

interface FileListVO {
  name: string
  file: string
  url: string
}

export default defineComponent({
  setup(prop, context) {
    const $axios: any = inject('$axios')
    let id: number = utils.getArticleId()
    let action = `${$axios.defaults.baseURL}/admin/files`
    let fileList = ref<FileListVO[]>([])
    onMounted(() => {
      $axios.get('/admin/files', {
        responseType: "json",
        params: {
          aid: id
        },
      }).then((response: any) => {
        response.data.forEach((v: FileList) => {
          let file: FileListVO = {
            name: v.name,
            file: v.file,
            url: utils.getUrl(v.file)
          }
          fileList.value.push(file)
        })
      }).catch((error: any) => {
        console.log(error)
      })
    })

    let removeFile = (file: string) => {
      $axios.delete('/admin/files', {
        params: {media: file},
        responseType: "json",
      }).then(() => {
        fileList.value = fileList.value.filter((v: FileListVO) => {
          return v.file != file
        })
      }).catch((error: any) => {
        ElMessage.warning(error.response.data.message)
      })
    }

    return {
      removeFile,
      action,
      fileList,
      insertDoc(file: string, name: string) {
        context.emit("insertDoc", name, file, utils.getUrl(file))
      },
      handleSuccess(response: any, file: any) {
        let data = new FormData()
        data.append("aid", id.toString())
        data.append("media", file.response.file_media)
        data.append("name", file.response.file_name)

        $axios.post('/admin/files/article', data, {
          responseType: "json",
        }).then(() => {
          file.file = file.response.file_media
          file.url = utils.getUrl(file.response.file_media)
          let t: FileListVO = {
            name: file.name,
            file: file.file,
            url: file.url,
          }
          fileList.value.push(t)
        }).catch((error: any) => {
          if (typeof error.response == 'string') {
            ElMessage.warning(error.response.data.message)
          } else {
            ElMessage.warning('系统异常，关系建立失败')
          }
          removeFile(file.file)
        })
      }
    }
  },
})
</script>
