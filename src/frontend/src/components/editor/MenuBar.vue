<template>
  <el-dialog
      v-model="fileUpload.show"
      width="30%"
      title="文件管理">
    <upload @insertDoc="file" />
  </el-dialog>
  <div>
    <right-menu @trigger="rightMenuTrigger"/>
    <template v-for="(item, index) in items">
      <div class="divider" v-if="item.type === 'divider'" :key="index"/>
      <menu-item :key="item.title" v-else v-bind="item"/>
    </template>
    <div class="menu-opt">
      <el-switch
          class="menu-opt-item"
          v-model="switchStatus.publish"
          active-text="公开"
          @change="save"
          :loading="switchLoading.publishLoading"></el-switch>
      <el-switch
          class="menu-opt-item"
          v-model="switchStatus.encrypt"
          active-text="加密"
          @change="save"
          :loading="switchLoading.encryptLoading"
          :beforeChange="encryptBeforeChange"></el-switch>
      <el-select
          @change="changeTopic"
          class="menu-opt-item"
          v-model="selectValue.topic"
          multiple
          allow-create
          filterable
          default-first-option
          size="mini"
          placeholder="选择专题">
        <el-option
            v-right-click="rightMenuFunc('topic', item.value)"
            v-for="item in topicOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value">
        </el-option>
      </el-select>
      <el-select
          @change="changeTag"
          class="menu-opt-item"
          v-model="selectValue.tag"
          multiple
          allow-create
          filterable
          default-first-option
          size="mini"
          placeholder="选择标签">
        <el-option
            v-right-click="rightMenuFunc('tag', item.value)"
            v-for="item in tagOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value">
        </el-option>
      </el-select>
    </div>
  </div>
</template>

<script lang="ts">
import '@fortawesome/fontawesome-free/css/all.css'
import MenuItem from './MenuItem.vue'
import {defineComponent, inject, provide, reactive, ref, toRefs} from 'vue'
import {ElMessage, ElMessageBox} from 'element-plus'
import utils from "@/common/utils";
import Upload from "@/components/editor/Upload.vue";

// eslint-disable-next-line
const enum CATEGORY_TYPE {
// eslint-disable-next-line
  Topic = "topic",
// eslint-disable-next-line
  Tag = "tag",
}

export default defineComponent({
  components: {
    MenuItem,
    Upload
  },

  emits: ["save"],

  props: {
    editor: {
      type: Object,
      required: true,
    },
  },

  setup(prop, context) {
    const $axios: any = inject('$axios')
    let switchStatus = reactive({
      encrypt: ref(false),
      publish: ref(false),
    })
    let switchLoading = reactive({
      publishLoading: ref(false),
      encryptLoading: ref(false),
    })
    switchStatus = inject('switch-status', switchStatus)

    // 用于检测是否设置过加密秘钥
    const encryptBeforeChange = () => {
      switchLoading.encryptLoading = true
      return new Promise(resolve => {
        // setTimeout(() => {
        switchLoading.encryptLoading = false
        ElMessage.success('切换成功')
        return resolve(true)
        // }, 500)
      })
    }

    let selectStatus = reactive({
      topicOptions: ref([{value: 0, label: ""}]),
      tagOptions: ref([{value: 0, label: ""}]),
    })

    let selectValue = reactive({
      topic: ref<number[]>([]),
      tag: ref<number[]>([]),
      topicOld: ref<number[]>([]),
      tagOld: ref<number[]>([]),
    })
    selectValue = inject("select-value", selectValue)


    // 查询专题列表
    $axios('/admin/category/topics', {
      method: "get",
      responseType: "json",
    }).then((response: any) => {
      selectStatus.topicOptions = []
      response.data.forEach((v: any) => {
        selectStatus.topicOptions.push({value: v.id, label: v.name})
      })
    }).catch((error: any) => {
      console.log(error)
    })

    // 查询标签列表
    $axios('/admin/category/tags', {
      method: "get",
      responseType: "json",
    }).then((response: any) => {
      selectStatus.tagOptions = []
      response.data.forEach((v: any) => {
        selectStatus.tagOptions.push({value: v.id, label: v.name})
      })
    }).catch((error: any) => {
      console.log(error)
    })

    // 自定义右键菜单
    let rightMenu = reactive({
      id: ref(0),
      position: ref({x: 0, y: 0}),
      show: ref(false),
      active: ref('article'),
      menu: ref({
        'topic': [
          {
            'icon': 'el-icon-edit',
            'title': '编辑',
            'command': 'topic_edit',
          },
          {
            'icon': 'el-icon-delete',
            'title': '删除',
            'command': 'topic_delete',
          },
        ],
        'tag': [
          {
            'icon': 'el-icon-edit',
            'title': '编辑',
            'command': 'tag_edit',
          },
          {
            'icon': 'el-icon-delete',
            'title': '删除',
            'command': 'tag_delete',
          },
        ]
      })
    })
    provide('right-click-menu', rightMenu)

    // 设置右键菜单点击参入值
    const rightMenuFunc = (active: string, id: number) => {
      return () => {
        rightMenu.active = active
        rightMenu.id = id
        return rightMenu
      }
    }

    // 执行移除分类操作
    const removeCategory = (type: CATEGORY_TYPE, id: number) => {
      $axios.delete(`/admin/category/${type}`, {
        params: {id: id},
        responseType: "json",
      }).then(() => {
        if (type == CATEGORY_TYPE.Tag) {
          selectValue.tagOld = selectValue.tag = selectValue.tag.filter((v: any) => {
            return v != id
          })
          selectStatus.tagOptions = selectStatus.tagOptions.filter((v: any) => {
            return v.value != id
          })
        } else {
          selectValue.topicOld = selectValue.topic = selectValue.topic.filter((v: any) => {
            return v != id
          })
          selectStatus.topicOptions = selectStatus.topicOptions.filter((v: any) => {
            return v.value != id
          })
        }
        ElMessage({
          type: 'success',
          message: '移除成功!',
        });
      }).catch((error: any) => {
        ElMessage.warning(error.response.data.message)
      })
    }

    // 确认移除弹窗
    const removeConfirm = (type: CATEGORY_TYPE, id: number) => {
      ElMessageBox.confirm('此操作将永久移除该分类, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
          .then(() => {
            removeCategory(type, id)
          })
          .catch(() => {
            ElMessage({
              type: 'info',
              message: '已取消移除',
            });
          });
    }

    // 更新分类名称操作
    const updateCategory = (type: CATEGORY_TYPE, id: number, name: string) => {
      let formData = new FormData();
      formData.append("name", name);
      formData.append("id", id.toString());
      $axios.put(`/admin/category/${type}`, formData, {
        headers: {'Content-Type': 'application/x-www-form-urlencoded'},
        responseType: "json",
      }).then(() => {
        if (type == CATEGORY_TYPE.Tag) {
          selectStatus.tagOptions = selectStatus.tagOptions.map((v: any) => {
            if (v.value == id) {
              v.label = name
            }
            return v
          })
        } else {
          selectStatus.topicOptions = selectStatus.topicOptions.map((v: any) => {
            if (v.value == id) {
              v.label = name
            }
            return v
          })
        }
        ElMessage({
          type: 'success',
          message: '更新成功!',
        });
      }).catch((error: any) => {
        ElMessage.warning(error.response.data.message)
      })
    }

    // 获取分类名
    const getLabel = (type: CATEGORY_TYPE, id: number): string => {
      let options: any
      if (type == CATEGORY_TYPE.Tag) {
        options = selectStatus.tagOptions
      } else {
        options = selectStatus.topicOptions
      }
      options = options.filter((v: any) => {
        return v.value == id
      })
      return options[0].label
    }

    // 分类编辑弹窗
    const editDialog = (type: CATEGORY_TYPE, id: number) => {
      ElMessageBox.prompt('请输入将要变更的分类名称', '编辑分类', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        inputValue: getLabel(type, id),
      }).then(({value}) => {
        updateCategory(type, id, value)
      })
          .catch(() => {
            ElMessage({
              type: 'info',
              message: '取消输入',
            });
          });
    }

    // 右键菜单逻辑处理
    const rightMenuTrigger = (type: string) => {
      switch (type) {
        case 'topic_delete':
          removeConfirm(CATEGORY_TYPE.Topic, rightMenu.id)
          break;
        case 'tag_delete':
          removeConfirm(CATEGORY_TYPE.Tag, rightMenu.id)
          break;
        case 'topic_edit':
          editDialog(CATEGORY_TYPE.Topic, rightMenu.id)
          break;
        case 'tag_edit':
          editDialog(CATEGORY_TYPE.Tag, rightMenu.id)
          break;
        default:
          break;
      }
    }

    // 新增分类名（专题/标签）
    const saveCategory = (type: CATEGORY_TYPE, name: string) => {
      $axios(`/admin/category/${type}`, {
        method: "post",
        responseType: "json",
        params: {name: name},
      }).then((response: any) => {
        if (type == CATEGORY_TYPE.Tag) {
          selectStatus.tagOptions.push({value: response.data, label: name})
        } else {
          selectStatus.topicOptions.push({value: response.data, label: name})
        }
        bindCategory(type, response.data)
      }).catch((error: any) => {
        ElMessage.warning(error.response.data.message)
      })
    }

    // 绑定分类（专题/标签）
    const bindCategory = (type: CATEGORY_TYPE, id: number) => {
      let aid = utils.getArticleId()
      if (aid == 0) {
        ElMessage.warning('绑定文章失败，需要先保存文章')
        return
      }
      $axios('/admin/category/bind' + type, {
        method: "post",
        responseType: "json",
        params: {id: id, aid: aid},
      }).then(() => {
        if (type == CATEGORY_TYPE.Tag) {
          selectValue.tag.push(id)
          selectValue.tagOld = selectValue.tag
        } else {
          selectValue.topic.push(id)
          selectValue.topicOld = selectValue.topic
        }
      }).catch((error: any) => {
        ElMessage.warning(error.response.data.message)
      })
    }

    // 解绑分类（专题/标签）
    const unbindCategory = (type: CATEGORY_TYPE, id: number) => {
      let aid = utils.getArticleId()
      if (aid == 0) {
        ElMessage.warning('解除绑定文章失败，需要先保存文章')
        return
      }
      $axios.delete('/admin/category/relations', {
        params: {id: id, aid: aid},
        responseType: "json",
      }).then(() => {
        if (type == CATEGORY_TYPE.Tag) {
          selectValue.tagOld = selectValue.tag = selectValue.tag.filter((v: any) => {
            return v != id
          })
        } else {
          selectValue.topicOld = selectValue.topic = selectValue.topic.filter((v: any) => {
            return v != id
          })
        }
      }).catch((error: any) => {
        ElMessage.warning(error.response.data.message)
      })
    }

    // 专题/标签发生变更时的处理逻辑（新增、绑定、解绑）
    const changeCore = (v: any, old: any, type: CATEGORY_TYPE) => {
      let os = new Set(old)
      let ns = new Set(v)

      if (type == CATEGORY_TYPE.Tag) {
        selectValue.tag = old
      } else {
        selectValue.topic = old
      }

      // 绑定
      if (old.length < v.length) {
        let val = v.filter((value: any) => {
          return !os.has(value)
        })
        if (isNaN(val[0])) {
          saveCategory(type, val[0])
        } else {
          bindCategory(type, val[0])
        }
      } else {
        // 解绑
        let val: number[] = old.filter((value: any) => {
          return !ns.has(value)
        })
        unbindCategory(type, val[0])
      }
    }

    // 标签发生变更时触发
    const changeTag = (v: any) => {
      changeCore(v, selectValue.tagOld, CATEGORY_TYPE.Tag)
    }

    // 专题发生变更时触发
    const changeTopic = (v: any) => {
      changeCore(v, selectValue.topicOld, CATEGORY_TYPE.Topic)
    }

    // 状态按钮发生变更时执行保存
    const save = () => {
      context.emit("save")
    }

    // 文件
    let fileUpload = reactive({show: ref<boolean>(false)})

    let file = (name: string, file: string, url: string) => {
      // console.log(name, file, url)
      if (file.search(/jpg|jpeg|gif|png|bmp/i) > -1) {
        prop.editor.chain().focus().setImage({src: url, title: name, alt: name}).run()
      } else {
        prop.editor.chain().focus().insertContent(`<a href="${url}" target="_blank">${name}</a>`).run()
      }
      // prop.editor
      //     .chain()
      //     .focus()
      //     .extendMarkRange('link')
      //     .setLink({ href: url })
      //     .run()
    }

    const items = [
        {
          icon: 'fas fa-bold',
          title: 'Bold',
          action: () => prop.editor.chain().focus().toggleBold().run(),
          isActive: () => prop.editor.isActive('bold'),
        },
        {
          icon: 'fas fa-italic',
          title: 'Italic',
          action: () => prop.editor.chain().focus().toggleItalic().run(),
          isActive: () => prop.editor.isActive('italic'),
        },
        {
          icon: 'fas fa-strikethrough',
          title: 'Strike',
          action: () => prop.editor.chain().focus().toggleStrike().run(),
          isActive: () => prop.editor.isActive('strike'),
        },
        {
          icon: 'fas fa-underline',
          title: 'Underline',
          action: () => prop.editor.chain().focus().toggleUnderline().run(),
          isActive: () => prop.editor.isActive('underline'),
        },
        {
          icon: 'fas fa-subscript',
          title: 'Subscript',
          action: () => prop.editor.chain().focus().toggleSubscript().run(),
          isActive: () => prop.editor.isActive('subscript'),
        },
        {
          icon: 'fas fa-superscript',
          title: 'Superscript',
          action: () => prop.editor.chain().focus().toggleSuperscript().run(),
          isActive: () => prop.editor.isActive('superscript'),
        },
        {
          icon: 'fas fa-lightbulb',
          title: 'Highlight',
          action: () => prop.editor.chain().focus().toggleHighlight().run(),
          isActive: () => prop.editor.isActive('highlight'),
        },
        {
          type: 'divider',
        },
        {
          icon: 'fas fa-align-left',
          title: 'left',
          action: () => prop.editor.chain().focus().setTextAlign('left').run(),
          isActive: () => prop.editor.isActive({textAlign: 'left'}),
        },
        {
          icon: 'fas fa-align-center',
          title: 'left',
          action: () => prop.editor.chain().focus().setTextAlign('center').run(),
          isActive: () => prop.editor.isActive({textAlign: 'center'}),
        },
        {
          icon: 'fas fa-align-right',
          title: 'right',
          action: () => prop.editor.chain().focus().setTextAlign('right').run(),
          isActive: () => prop.editor.isActive({textAlign: 'right'}),
        },
        {
          icon: 'fas fa-align-justify',
          title: 'justify',
          action: () => prop.editor.chain().focus().setTextAlign('justify').run(),
          isActive: () => prop.editor.isActive({textAlign: 'justify'}),
        },
        {
          type: 'divider',
        },
        {
          icon: 'fas fa-heading',
          title: 'Heading',
          action: () => prop.editor.chain().focus().toggleHeading({level: 2}).run(),
          isActive: () => prop.editor.isActive('heading', {level: 2}),
        },
        {
          icon: 'fas fa-paragraph',
          title: 'Paragraph',
          action: () => prop.editor.chain().focus().setParagraph().run(),
          isActive: () => prop.editor.isActive('paragraph'),
        },
        {
          icon: 'fas fa-list-ul',
          title: 'Bullet List',
          action: () => prop.editor.chain().focus().toggleBulletList().run(),
          isActive: () => prop.editor.isActive('bulletList'),
        },
        {
          icon: 'fas fa-list-ol',
          title: 'Ordered List',
          action: () => prop.editor.chain().focus().toggleOrderedList().run(),
          isActive: () => prop.editor.isActive('orderedList'),
        },
        {
          icon: 'fas fa-tasks',
          title: 'Task List',
          action: () => prop.editor.chain().focus().toggleTaskList().run(),
          isActive: () => prop.editor.isActive('taskList'),
        },
        {
          icon: 'fas fa-terminal',
          title: 'Code Block',
          action: () => prop.editor.chain().focus().toggleCodeBlock().run(),
          isActive: () => prop.editor.isActive('codeBlock'),
        },
        {
          type: 'divider',
        },
        {
          icon: 'fas fa-quote-right',
          title: 'Blockquote',
          action: () => prop.editor.chain().focus().toggleBlockquote().run(),
          isActive: () => prop.editor.isActive('blockquote'),
        },
        {
          icon: 'fas fa-minus',
          title: 'Horizontal Rule',
          action: () => prop.editor.chain().focus().setHorizontalRule().run(),
        },
      {
        icon: 'fas fa-upload',
        title: 'Upload',
        action: () => {
          if (utils.getArticleId() < 1) {
            save()
            setTimeout(() => {
              if (utils.getArticleId() < 1) {
                ElMessage.warning("保存文章后再试")
                return
              }
              fileUpload.show  = true
            }, 2000)
          } else {
            fileUpload.show  = true
          }
        },
      },
      {
        type: 'divider',
      },
      {
        icon: 'fas fa-level-down-alt',
        title: 'Hard Break',
        action: () => prop.editor.chain().focus().setHardBreak().run(),
      },
      {
        icon: 'fas fa-eraser',
        title: 'Clear Format',
        action: () => prop.editor.chain()
            .focus()
            .clearNodes()
            .unsetAllMarks()
            .run(),
      },
      {
        type: 'divider',
      },
      {
        icon: 'fas fa-undo',
        title: 'Undo',
        action: () => prop.editor.chain().focus().undo().run(),
      },
      {
        icon: 'fas fa-redo',
        title: 'Redo',
        action: () => prop.editor.chain().focus().redo().run(),
      },
    ]

    return {
      file,
      save,
      items,
      fileUpload,
      changeTag,
      changeTopic,
      switchStatus,
      selectValue,
      switchLoading,
      ...toRefs(selectStatus),
      encryptBeforeChange,
      rightMenu,
      rightMenuFunc,
      rightMenuTrigger,
    }

  },
})
</script>

<style lang="scss" scoped>
.divider {
  width: 1px;
  margin-left: 1rem;
  display: inline-block;
}

.menu-opt {
  padding-top: 0.5rem;
  line-height: 2rem;

.menu-opt-item {
  margin-right: 12px;
}

}
</style>
