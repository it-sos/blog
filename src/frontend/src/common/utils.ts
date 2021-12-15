import router from "../routes";
import axios from "axios";

export default {
   getArticleId: () => {
      if (typeof router.currentRoute.value.params.id != "undefined") {
         return parseInt(router.currentRoute.value.params.id.toString())
      }
      return 0
   },
   getUrl: (media: string) => {
      media = media.replace("/", "_")
      return `${axios.defaults.baseURL}/files/${media}`
   }
}
