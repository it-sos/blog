import { BubbleMenu } from "@tiptap/extension-bubble-menu";
import CodeBlockLowlight from "@tiptap/extension-code-block-lowlight";
import Highlight from "@tiptap/extension-highlight";
import Image from "@tiptap/extension-image";
import { Link } from "@tiptap/extension-link";
import Subscript from "@tiptap/extension-subscript";
import Superscript from "@tiptap/extension-superscript";
import TaskItem from "@tiptap/extension-task-item";
import TaskList from "@tiptap/extension-task-list";
import TextAlign from "@tiptap/extension-text-align";
import Typography from "@tiptap/extension-typography";
import Underline from "@tiptap/extension-underline";
import StarterKit from "@tiptap/starter-kit";
import { Extensions, VueNodeViewRenderer } from "@tiptap/vue-3";
import lowlight from "lowlight";
import CodeBlockComponent from "../../components/editor/CodeBlockComponent.vue";
import PasteHandle from './paste-handle';

const CustomCodeBlock = CodeBlockLowlight
    .extend({
        addNodeView() {
            return VueNodeViewRenderer(CodeBlockComponent)
        },
    })
    .configure({ lowlight })

const common = [
    // Document,
    // Paragraph,
    Link,
    // Text,
    Image,
    Underline,
    Subscript,
    Superscript,
    TaskList,
    TaskItem,
    StarterKit,
    TextAlign.configure({
        types: ['heading', 'paragraph'],
    }),
    Highlight,
    Typography,
]

export const backendExtensions: Extensions = [
    ...common,
    CustomCodeBlock,
    BubbleMenu.configure({
        shouldShow: ({ editor, view, state, oldState, from, to }) => {
            // only show the bubble menu for images and links
            return editor.isActive('image')
        },
    }),
    PasteHandle,
];

export const frontendExtensions: Extensions = [
    CodeBlockLowlight.configure({ lowlight }),
    ...common,
]