import {Extension, Node} from '@tiptap/core'
import {Plugin} from 'prosemirror-state'
// @ts-ignore
import {serializeForClipboard} from 'prosemirror-view/src/clipboard'
import axios from "axios";
import {ElMessage} from "element-plus";
import utils from "@/common/utils";

// @ts-ignore
export default Extension.create({
    addProseMirrorPlugins() {
        return [
            new Plugin({
                props: {
                    handlePaste(view: any, event: any): boolean {
                        let hasFiles =
                            event.clipboardData &&
                            event.clipboardData.files &&
                            event.clipboardData.files.length;

                        if (!hasFiles) {
                            return false;
                        }

                        const removeFile = (file: string) => {
                            axios.delete('/admin/files', {
                                params: {media: file},
                                responseType: "json",
                            }).then(() => {
                            }).catch((error: any) => {
                                ElMessage.warning(error.response.data.message)
                            })
                        }

                        // 从剪贴板中读取文件
                        Array.from(event.clipboardData.files)
                            .forEach((item: any) => {//读取数据
                                const data = new FormData()
                                data.append("file", item)
                                axios.post('/admin/files', data, {
                                    responseType: "json",
                                }).then((response: any) => {
                                    const data = new FormData()
                                    data.append("aid", utils.getArticleId().toString())
                                    data.append("media", response.data.file_media)
                                    data.append("name", response.data.file_name)

                                    axios.post('/admin/files/article', data, {
                                        responseType: "json",
                                    }).then(() => {
                                        if (response.data.file_name.search(/jpg|jpeg|gif|png|bmp/i) > -1) {
                                            // 以图像方式展示
                                            const node = view.state.schema.nodes.image.create({src: utils.getUrl(response.data.file_media), alt: response.data.file_name, title: response.data.file_name});
                                            const transaction = view.state.tr.replaceSelectionWith(node);
                                            view.dispatch(transaction);
                                        } else {
                                            // 非图像以链接方式展示
                                            const href = utils.getUrl(response.data.file_media)
                                            const text = response.data.file_name
                                            const mark = view.state.schema.marks.link.create({ href })
                                            const from = view.state.selection.from
                                            const transaction = view.state.tr.insertText(text + ' ')
                                            transaction.addMark(from, from + text.length, mark)
                                            view.dispatch(transaction);
                                        }
                                    }).catch((error: any) => {
                                        if (typeof error.response == 'string') {
                                            ElMessage.warning(error.response.data.message)
                                        } else {
                                            ElMessage.warning('系统异常，关系建立失败')
                                        }
                                        removeFile(response.data.file_media)
                                    })
                                }).catch((error: any) => {
                                    if (typeof error.response == 'string') {
                                        ElMessage.warning(error.response.data.message)
                                    } else {
                                        ElMessage.warning('系统异常，稍後重試')
                                    }
                                })
                                hasFiles = true;
                            });
                        //扫尾
                        if (hasFiles) {
                            event.preventDefault();
                            return true;
                        }
                        return false;
                    },
                },
            }),
        ]
    },
})
