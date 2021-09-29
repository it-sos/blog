<template>
  <el-upload
      class="upload-demo"
      :action="action"
      :on-preview="handlePreview"
      :on-remove="handleRemove"
      :file-list="fileList"
      :on-success="handleSuccess"
      multiple>
    <el-button size="small" type="primary">Click to upload</el-button>
    <template #tip>
      <div class="el-upload__tip">
        files with a size less than 500kb
      </div>
    </template>
  </el-upload>
</template>
<script lang="ts">
import {defineComponent, onMounted, ref} from "vue";
import axios from "axios";
import utils from "@/common/utils";
import {ElMessage} from "element-plus";


interface FileList {
  name?: string
  file?: string
}

interface FileListVO {
  name?: string
  url?: string
}

export default defineComponent({
  setup() {
    let id: number = utils.getArticleId()
    let action = `${axios.defaults.baseURL}/admin/files`
    let fileList = ref<FileListVO[]>([])
    onMounted(() => {
      axios.get('/admin/files', {
        responseType: "json",
        params: {
          aid: id
        },
      }).then((response: any) => {
        response.data.forEach((v: FileList) => {
          let file: FileListVO = {
            name: v.name,
            url: `/file?media=${v.file}`
          }
          fileList.value.push(file)
        })
      }).catch((error: any) => {
        console.log(error)
      })
    })

    return {
      action,
      fileList,
      handleRemove(file: any) {
        axios.delete('/admin/files', {
          params: {media: file.response.file_media},
          responseType: "json",
        }).then((response: any) => {
          console.log(response)
        }).catch((error: any) => {
          ElMessage.warning(error.response.data.message)
        })
      },
      handlePreview(file: any) {
        console.log(file)
      },
      handleSuccess(response: any, file: any) {
        let data = new FormData()
        data.append("aid", id.toString())
        data.append("media", file.response.file_media)
        data.append("name", file.response.file_name)
        axios.post('/admin/files/article', data, {
          responseType: "json",
        }).then((response: any) => {
          console.log(response)
        }).catch((error: any) => {
          ElMessage.warning(error.response.data.message)
        })
      }
    }
  },
})
</script>
