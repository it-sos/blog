<template>
  <el-upload
      class="upload-demo"
      action="https://jsonplaceholder.typicode.com/posts/"
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
import {defineComponent} from "vue";

export default defineComponent({
  setup() {
    return {
      fileList: [
        {
          name: 'food.pdf',
          url: 'http://baid.com/pdf',
        },
        {
          name: 'food2.jpeg',
          url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100',
        },
      ],
      handleRemove(file: any, fileList: any) {
        console.log(file, fileList)
      },
      handlePreview(file: any) {
        console.log(file)
      },
      handleSuccess(response: any, file: any, fileList: any) {
        console.log(response, file, fileList)
      }
    }
  },
})
</script>
