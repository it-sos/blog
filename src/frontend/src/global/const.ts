const SearchSite = '1'
const SearchAI = '2'

// 文章加密状态
// eslint-disable-next-line
const enum SWITCH_ENCRYPT_STATUS {
    // 密文
    // eslint-disable-next-line
    Encrypt = 1,
    // 明文
    // eslint-disable-next-line
    Plaintext = 2,
}

// 文章发布状态
// eslint-disable-next-line
const enum SWITCH_PUBLISH_STATUS {
    // 私有
    // eslint-disable-next-line
    Private = 1,
    // 公开
    // eslint-disable-next-line
    Public = 2,
}

export { SWITCH_ENCRYPT_STATUS, SWITCH_PUBLISH_STATUS, SearchSite, SearchAI }
