import {Extensions, VueNodeViewRenderer} from "@tiptap/vue-3";
import Document from "@tiptap/extension-document";
import Paragraph from "@tiptap/extension-paragraph";
import Text from "@tiptap/extension-text";
import {Link} from "@tiptap/extension-link";
import Underline from "@tiptap/extension-underline";
import Subscript from "@tiptap/extension-subscript";
import Superscript from "@tiptap/extension-superscript";
import TaskList from "@tiptap/extension-task-list";
import TaskItem from "@tiptap/extension-task-item";
import StarterKit from "@tiptap/starter-kit";
import TextAlign from "@tiptap/extension-text-align";
import Highlight from "@tiptap/extension-highlight";
import Typography from "@tiptap/extension-typography";
import CodeBlockLowlight from "@tiptap/extension-code-block-lowlight";
import CodeBlockComponent from "../../components/editor/CodeBlockComponent.vue";
import Image from "@tiptap/extension-image";
import lowlight from "lowlight";
import {BubbleMenu} from "@tiptap/extension-bubble-menu";
import PasteHandle from './paste-handle'

const CustomCodeBlock = CodeBlockLowlight
    .extend({
        addNodeView() {
            return VueNodeViewRenderer(CodeBlockComponent)
        },
    })
    .configure({lowlight})

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
        shouldShow: ({editor, view, state, oldState, from, to}) => {
            // only show the bubble menu for images and links
            return editor.isActive('image')
        },
    }),
    PasteHandle,
];

export const frontendExtensions: Extensions = [
    CodeBlockLowlight.configure({lowlight}),
    ...common
]