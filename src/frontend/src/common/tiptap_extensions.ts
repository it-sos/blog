import {Extensions, VueNodeViewRenderer} from "@tiptap/vue-3";
import Document from "@tiptap/extension-document";
import Paragraph from "@tiptap/extension-paragraph";
import Text from "@tiptap/extension-text";
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
import CodeBlockComponent from "@/components/editor/CodeBlockComponent.vue";
import lowlight from "lowlight";

let CodeBlock = CodeBlockLowlight
    .extend({
        addNodeView() {
            return VueNodeViewRenderer(CodeBlockComponent)
        },
    })
    .configure({lowlight})

const extensions: Extensions = [
    Document,
    Paragraph,
    Text,
    Underline,
    Subscript,
    Superscript,
    CodeBlock,
    TaskList,
    TaskItem,
    StarterKit,
    TextAlign.configure({
        types: ['heading', 'paragraph'],
    }),
    Highlight,
    Typography,
];
export default extensions