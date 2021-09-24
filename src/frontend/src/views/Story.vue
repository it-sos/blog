<template>
  <el-container>
    <el-aside width="200px" class="hidden-xs-only">
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <el-input
                placeholder="请输入内容"
                v-model.trim="keyword"
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
          <div v-bind:key="art.id" v-for="art in article" class="text item dirs">
            <el-link href="javascript:void(0);" @click="edit(art.id)"
                     v-right-click="rightMenuFunc('article', art.id)"><span
                v-html="art.title_match ? art.title_match : art.title"></span></el-link>
            <el-tag style="margin-left:0.3rem;" effect="plain" type="danger" size="mini">{{ art.duration }}</el-tag>
          </div>
        </div>
      </el-card>
    </el-aside>
    <el-main>
      <tool-editor @syncArticleList="syncArticleList"/>
    </el-main>
  </el-container>
  <right-menu @trigger="rightMenuTrigger"/>
</template>
<script lang="ts">

import {defineComponent, provide, reactive, ref, toRefs, watch} from 'vue'
import {ElMessage, ElMessageBox} from "element-plus";
import axios from "axios";
import {router} from "@/routes";
import utils from "@/common/utils";
import ToolEditor from "@/components/ToolEditor.vue";

interface ArticleList {
  id?: number
  title: string
  title_match?: string
  duration?: string
}

export default defineComponent({
  components: {
    ToolEditor,
  },

  setup() {
    document.title = "editing"

    const state = reactive({
      keyword: ref(""),
      page: ref(0),
      size: ref(10),
      noMore: ref(false),
      article: ref<ArticleList[]>([]),
    });

    let timer: any = null
    watch(() => state.keyword, () => {
      clearTimeout(timer)
      timer = setTimeout(articleListRest, 1000)
    })

    const load = () => {
      if (state.noMore) {
        return
      }
      state.page++
      axios.get('/admin/articles', {
        responseType: "json",
        params: {page: state.page, size: state.size, keyword: state.keyword},
      }).then((response) => {
        response.data.forEach((v: ArticleList) => {
          state.article.push(v)
        })
        if (response.data.length < state.size) {
          state.noMore = true
        }
      }).catch((error: any) => {
        ElMessage.warning(error.response.data.message)
      })
    }

    // 重置文章列表
    const articleListRest = () => {
      state.page = 0
      state.noMore = false
      state.article = []
      load()
    }

    // eslint-disable-next-line
    const enum OPT_TYPE {
      // eslint-disable-next-line
      Add = "add",
      // eslint-disable-next-line
      Update = "update"
    }

    // 编辑文章并动态同步文章列表
    const syncArticleList = (type: OPT_TYPE, id: number, title: string) => {
      // 更新
      if (type == OPT_TYPE.Update) {
        state.article = state.article.map((v: ArticleList) => {
          if (v.id == id) {
            v.title = title
          }
          return v
        })
      }
      // 新增
      else if (type == OPT_TYPE.Add) {
        state.article.unshift({
          id: id,
          title: title,
          duration: '1秒前',
        })
      }
    }

    // 删除操作
    const deleteConfirm = (id: number) => {
      ElMessageBox.confirm('此操作将永久删除该文章, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      }).then(() => {
        axios.delete('/admin/article', {
          responseType: "json",
          params: {id: id},
        }).then(() => {
          state.article = state.article.filter((v: any) => {
            return id != v.id
          })
          ElMessage({
            type: 'success',
            message: '删除成功!',
          });

          if (utils.getArticleId() == id) {
            router.push('/e/')
          }
        }).catch((error: any) => {
          ElMessage.warning(error.response.data.message)
        })
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
          // {
          //   'icon': 'el-icon-edit',
          //   'title': '编辑',
          //   'command': 'article_edit',
          // },
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
      switch (type) {
        case 'article_add':
          router.push('/e/')
          break;
          // case 'article_edit':
          //   router.push('/e/'+rightMenu.id)
          //   break;
        case 'article_delete':
          deleteConfirm(rightMenu.id)
          break;
      }
    }
    const edit = (id: number) => {
      router.push('/e/' + id)
    }

    return {
      ...toRefs(state),
      edit,
      load,
      syncArticleList,
      rightMenu,
      rightMenuFunc,
      rightMenuTrigger,
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
