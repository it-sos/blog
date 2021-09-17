import {router} from "@/routes";

export default {
   getArticleId: () => {
      if (typeof router.currentRoute.value.params.id != "undefined") {
         return parseInt(router.currentRoute.value.params.id.toString())
      }
      return 0
   }
}