<template>
  <el-container>
    <el-aside width="200px" class="hidden-xs-only">
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <el-input
                placeholder="请输入内容"
                v-model="keyword"
                class="input-with-select"
            >
              <template #prefix>
                <i class="el-input__icon el-icon-search"></i>
              </template>
            </el-input>
          </div>
        </template>
        <div class="infinite-list-wrapper" style="height:400px;overflow-y:auto;overflow-x: hidden;"
             v-infinite-scroll="load">
          <div v-bind:key="idx" v-for="(art, idx) in article" class="text item dirs">
            <el-link href="javascript:void(0);" v-right-click="rightMenuFunc('article', idx)"><span
                v-html="art.article.title_match ? art.article.title_match : art.article.title"></span></el-link>
            <el-tag style="margin-left:0.3rem;" effect="plain" type="danger" size="mini">{{ art.duration }}</el-tag>
          </div>
        </div>
      </el-card>
    </el-aside>
    <el-main>
      <tool-editor/>
    </el-main>
  </el-container>
  <right-menu @trigger="rightMenuTrigger"/>
</template>
<script lang="ts">

import ToolEditor from '../components/ToolEditor.vue'

import {defineComponent, provide, reactive, ref, toRefs, watch} from 'vue'
import {ElMessage, ElMessageBox} from "element-plus";

export default defineComponent({
  components: {
    ToolEditor,
  },

  setup() {
    document.title = "editing"

    let article: any[] = []
    const state = reactive({
      keyword: ref(""),
      page: ref(0),
      size: ref(10),
      noMore: ref(false),
      article: article,
    });

    const defaults = () => {
      state.page = 0
      state.noMore = false
    }

    const open = () => {
      ElMessageBox.confirm('此操作将永久删除该文件, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }).then(() => {
        ElMessage({
          type: 'success',
          message: '删除成功!',
        });
      }).catch(() => {
        ElMessage({
          type: 'info',
          message: '已取消删除',
        });
      });
    };

    let rightMenu = reactive({
      id: ref(0),
      position: ref({x: 0, y: 0}),
      show: ref(false),
      active: ref('article'),
      menu: ref({
        'article': [
          {
            'icon': 'el-icon-document-add',
            'title': '新建',
            'command': 'article_add',
          },
          {
            'icon': 'el-icon-edit',
            'title': '编辑',
            'command': 'article_edit',
          },
          {
            'icon': 'el-icon-delete',
            'title': '删除',
            'command': 'article_delete',
          },
        ],
      })
    })
    provide('right-click-menu', rightMenu)

    const rightMenuFunc = (active: string, id: number) => {
      return () => {
        rightMenu.active = active
        rightMenu.id = id
        return rightMenu
      }
    }

    const rightMenuTrigger = (type: string) => {
      console.log(type)
      console.log(rightMenu.id)
      open()
    }

    return {
      ...toRefs(state),
      defaults,
      rightMenu,
      rightMenuFunc,
      rightMenuTrigger,
    }
  },

  created() {
    let timer: any = null
    watch(
        () => this.keyword,
        () => {
          clearTimeout(timer)
          timer = setTimeout(() => {
            this.defaults()
            this.load()
          }, 1000)
        }
    )
  },

  methods: {
    load(): void {
      if (this.noMore) {
        return
      }
      this.article = [
        {
          article: {
            title: "标准标题1"
          },
          duration: "2秒"
        },
        {
          article: {
            title: "标准标题2"
          },
          duration: "2秒"
        },
      ]
      // this.page++
      // this.$http.get('/article/list', {
      //   params: {
      //     "page": this.page,
      //     "size": this.size,
      //     // "keyword": this.keyword
      //   }
      // }).then((response) => {
      //   if (response.data.length === 0) {
      //     this.noMore = true
      //     return
      //   }
      //   response.data.forEach((v: any) => {
      //     this.article.push(v)
      //   })
      // }).catch((error) => {
      //   console.log(error)
      // })
    }
  }
})
</script>

<style lang="scss" scoped>
.hidden-xs-only {
  padding-top: 20px;
}

.dirs {
  margin-bottom: 8px;
  display: flex;
  justify-content: space-between;

span {
  font-size: 12px;
}

}

.el-tabs--border-card > .el-tabs__content {

.el-card {
  border: none;
  box-shadow: none;
}

}

</style>
<style lang="scss">
.user-pane > .el-tabs__content {
  padding: 0;
}
</style>
