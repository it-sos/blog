<template>
  <div>
    <right-menu @trigger="rightMenuTrigger"/>
    <template v-for="(item, index) in items">
      <div class="divider" v-if="item.type === 'divider'" :key="index"/>
      <menu-item :key="item.title" v-else v-bind="item"/>
    </template>
    <div class="menu-opt">
      <el-switch
          class="menu-opt-item"
          v-model="publish"
          active-text="公开"
          :loading="publishLoading"
          :beforeChange="publishBeforeChange"></el-switch>
      <el-switch
          class="menu-opt-item"
          v-model="encrypt"
          active-text="加密"
          :loading="encryptLoading"
          :beforeChange="encryptBeforeChange"></el-switch>
      <el-select
          @change="changeTopic"
          class="menu-opt-item"
          v-model="topicValue"
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
          v-model="tagValue"
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
import {defineComponent, provide, reactive, ref, toRefs} from 'vue'
import {ElMessage} from 'element-plus'

export default defineComponent({
  components: {
    MenuItem,
  },

  props: {
    editor: {
      type: Object,
      required: true,
    },
  },

  setup() {

    const publishStatus = reactive({
      publish: false,
      publishLoading: false,
    })

    const encryptStatus = reactive({
      encrypt: false,
      encryptLoading: false,
    })

    const publishBeforeChange = () => {
      publishStatus.publishLoading = true
      return new Promise(resolve => {
        setTimeout(() => {
          publishStatus.publishLoading = false
          ElMessage.success('切换成功')
          return resolve(true)
        }, 1000)
      })
    }

    const encryptBeforeChange = () => {
      encryptStatus.encryptLoading = true
      return new Promise(resolve => {
        setTimeout(() => {
          encryptStatus.encryptLoading = false
          ElMessage.success('切换成功')
          return resolve(true)
        }, 1000)
      })
    }

    const selectStatus = reactive({
      topicValue: ['选项1'],
      topicOptions: [{
        value: '选项1',
        label: 'mac 下以 root 角色开机启动执行脚本或命令'
      }, {
        value: '选项5',
        label: '北京烤鸭'
      }],
      tagValue: [],
      tagOptions: [{
        lable:"php",
        value:"php"
      }],
    })

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
    }

    const changeTag = () => {
      console.log(1111)
    }

    const changeTopic = () => {
      console.log(1111)
    }

    return {
      changeTag,
      changeTopic,
      ...toRefs(publishStatus),
      ...toRefs(encryptStatus),
      ...toRefs(selectStatus),
      publishBeforeChange,
      encryptBeforeChange,
      rightMenu,
      rightMenuFunc,
      rightMenuTrigger,
    }

  },

  data() {
    return {
      items: [
        {
          icon: 'fas fa-bold',
          title: 'Bold',
          action: () => this.editor.chain().focus().toggleBold().run(),
          isActive: () => this.editor.isActive('bold'),
        },
        {
          icon: 'fas fa-italic',
          title: 'Italic',
          action: () => this.editor.chain().focus().toggleItalic().run(),
          isActive: () => this.editor.isActive('italic'),
        },
        {
          icon: 'fas fa-strikethrough',
          title: 'Strike',
          action: () => this.editor.chain().focus().toggleStrike().run(),
          isActive: () => this.editor.isActive('strike'),
        },
        {
          icon: 'fas fa-underline',
          title: 'Underline',
          action: () => this.editor.chain().focus().toggleUnderline().run(),
          isActive: () => this.editor.isActive('underline'),
        },
        {
          icon: 'fas fa-subscript',
          title: 'Subscript',
          action: () => this.editor.chain().focus().toggleSubscript().run(),
          isActive: () => this.editor.isActive('subscript'),
        },
        {
          icon: 'fas fa-superscript',
          title: 'Superscript',
          action: () => this.editor.chain().focus().toggleSuperscript().run(),
          isActive: () => this.editor.isActive('superscript'),
        },
        {
          icon: 'fas fa-lightbulb',
          title: 'Highlight',
          action: () => this.editor.chain().focus().toggleHighlight().run(),
          isActive: () => this.editor.isActive('highlight'),
        },
        {
          type: 'divider',
        },
        {
          icon: 'fas fa-align-left',
          title: 'left',
          action: () => this.editor.chain().focus().setTextAlign('left').run(),
          isActive: () => this.editor.isActive({textAlign: 'left'}),
        },
        {
          icon: 'fas fa-align-center',
          title: 'left',
          action: () => this.editor.chain().focus().setTextAlign('center').run(),
          isActive: () => this.editor.isActive({textAlign: 'center'}),
        },
        {
          icon: 'fas fa-align-right',
          title: 'right',
          action: () => this.editor.chain().focus().setTextAlign('right').run(),
          isActive: () => this.editor.isActive({textAlign: 'right'}),
        },
        {
          icon: 'fas fa-align-justify',
          title: 'justify',
          action: () => this.editor.chain().focus().setTextAlign('justify').run(),
          isActive: () => this.editor.isActive({textAlign: 'justify'}),
        },
        {
          type: 'divider',
        },
        {
          icon: 'fas fa-heading',
          title: 'Heading',
          action: () => this.editor.chain().focus().toggleHeading({level: 2}).run(),
          isActive: () => this.editor.isActive('heading', {level: 2}),
        },
        {
          icon: 'fas fa-paragraph',
          title: 'Paragraph',
          action: () => this.editor.chain().focus().setParagraph().run(),
          isActive: () => this.editor.isActive('paragraph'),
        },
        {
          icon: 'fas fa-list-ul',
          title: 'Bullet List',
          action: () => this.editor.chain().focus().toggleBulletList().run(),
          isActive: () => this.editor.isActive('bulletList'),
        },
        {
          icon: 'fas fa-list-ol',
          title: 'Ordered List',
          action: () => this.editor.chain().focus().toggleOrderedList().run(),
          isActive: () => this.editor.isActive('orderedList'),
        },
        {
          icon: 'fas fa-tasks',
          title: 'Task List',
          action: () => this.editor.chain().focus().toggleTaskList().run(),
          isActive: () => this.editor.isActive('taskList'),
        },
        {
          icon: 'fas fa-terminal',
          title: 'Code Block',
          action: () => this.editor.chain().focus().toggleCodeBlock().run(),
          isActive: () => this.editor.isActive('codeBlock'),
        },
        {
          type: 'divider',
        },
        {
          icon: 'fas fa-quote-right',
          title: 'Blockquote',
          action: () => this.editor.chain().focus().toggleBlockquote().run(),
          isActive: () => this.editor.isActive('blockquote'),
        },
        {
          icon: 'fas fa-minus',
          title: 'Horizontal Rule',
          action: () => this.editor.chain().focus().setHorizontalRule().run(),
        },
        {
          type: 'divider',
        },
        {
          icon: 'fas fa-level-down-alt',
          title: 'Hard Break',
          action: () => this.editor.chain().focus().setHardBreak().run(),
        },
        {
          icon: 'fas fa-eraser',
          title: 'Clear Format',
          action: () => this.editor.chain()
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
          action: () => this.editor.chain().focus().undo().run(),
        },
        {
          icon: 'fas fa-redo',
          title: 'Redo',
          action: () => this.editor.chain().focus().redo().run(),
        },
      ],
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
