import {ElMessage, ElMessageBox} from "element-plus";
import utils from "../../common/utils";
import axios from "axios";
import router from "../../routes";

interface State {
    data: ArticleList[]
}

export default {
    namespaced: true,
    state: () => ({
        data: []
    }),
    mutations: {
        append(state: State, data: ArticleList) {
            state.data.push(data)
        },
        add(state: State, data: ArticleList) {
            state.data.unshift({
                id: data.id,
                title: data.title,
                duration: '1秒前',
                is_state: data.is_state,
            })
        },
        remove(state: State, data: ArticleList) {
            ElMessageBox.confirm('此操作将永久删除该文章, 是否继续?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning',
            }).then(() => {
                axios.delete('/admin/article', {
                    responseType: "json",
                    params: {id: data.id},
                }).then(() => {
                    state.data = state.data.filter((v: any) => {
                        return data.id != v.id
                    })
                    ElMessage({
                        type: 'success',
                        message: '删除成功!',
                    });

                    if (utils.getArticleId() == data.id) {
                        router.push('/e/')
                    }
                }).catch((error: any) => {
                    ElMessage.warning(error.response.data.message)
                })
            }).catch((e: any) => {
                ElMessage({
                    type: 'info',
                    message: '已取消删除',
                });
            });
        },
        update(state: State, data: ArticleList) {
            state.data = state.data.map((v: ArticleList) => {
                if (v.id == data.id) {
                    v.title = data.title
                    v.is_state = data.is_state
                }
                return v
            })
        },
        reset(state: State) {
            state.data = []
        },
    },
    getters: {
        list(state: State): ArticleList[] {
            return state.data
        }
    }
}
