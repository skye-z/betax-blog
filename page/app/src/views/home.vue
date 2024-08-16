<template>
    <div class="app-content no-select">
        <div class="flex mb-10">
            <article-banner class="mr-10" :list="banner" />
            <div class="full-width">
                <div class="card readme mb-10 full-width">
                    <div class="flex readme-info pa-10 border-bottom" v-if="user.nickname">
                        <n-avatar class="readme-avatar mr-10" size="large" :src="user.avatar" />
                        <div>
                            <div class="flex align-center">
                                <div class="mr-5">{{ user.nickname }}</div>
                                <div class="text-gray text-small">@{{ user.username }}</div>
                            </div>
                            <div class="text-gray text-small lint1">{{ user.bio }}</div>
                        </div>
                    </div>
                    <div class="bind-tips text-center border-bottom" v-else>
                        请点击下方按钮绑定 Github 账户
                    </div>
                    <n-button v-if="user.nickname" class="full-width" quaternary @click="openGithub">
                        <template #icon>
                            <n-icon>
                                <Github />
                            </n-icon>
                        </template>
                        前往 Github 主页
                    </n-button>
                    <a v-else href="/login" rel="console">
                        <n-button class="full-width" quaternary>
                            <template #icon>
                                <n-icon>
                                    <Github />
                                </n-icon>
                            </template>
                            绑定账户
                        </n-button>
                    </a>
                </div>
                <div class="card tag-cloud pt-10 pl-10 full-width">
                    <n-tag class="tag-item mr-10 mb-10" v-for="item in tags" :bordered="false" type="success" @click="openTag(item.id)">
                        {{ item.name }}
                    </n-tag>
                </div>
            </div>
        </div>
        <div class="article-box">
            <article-list :list="tops" />
            <article-list :list="list" />
            <div class="text-center">
                <n-button class="load-more" :disabled="!next" strong secondary
                    @click="getList(false, false)">加载更多</n-button>
            </div>
        </div>
    </div>
</template>

<script>
import ArticleList from '../components/articleList.vue'
import ArticleBanner from '../components/articleBanner.vue'
import { Github, Linkedin, Discord } from '@vicons/fa'
import { QqOutlined } from '@vicons/antd'
import { common, article } from '../plugins/api'

export default {
    name: "Home",
    components: { Github, Linkedin, QqOutlined, Discord, ArticleList, ArticleBanner },
    data: () => ({
        loading: true,
        next: true,
        number: 20,
        page: 1,
        list: [],
        tops: [],
        banner: [],
        tags: [],
        user: {}
    }),
    methods: {
        init() {
            let tags = localStorage.getItem('cache:tags');
            if (tags != undefined) {
                this.tags = JSON.parse(tags)
            }
            let user = localStorage.getItem('cache:user');
            if (user != undefined) {
                this.user = JSON.parse(user)
            }
            common.init().then(res => {
                if (res.state) {
                    this.banner = res.data.banner ? res.data.banner : []
                    this.tops = res.data.up ? res.data.up : []
                    this.buildList(res.data.list)
                    this.loading = false;
                } else {
                    window.$message.warning('初始化失败');
                    this.loading = false;
                }
            }).catch(err => {
                this.loading = false;
                window.$message.warning("初始化出错");
            })
        },
        getList(isBanner, isUp) {
            article.getList(isBanner, isUp, -1, -1, 2, this.page, this.number).then(res => {
                if (res.state) {
                    this.buildList(res.data)
                    this.loading = false;
                } else {
                    window.$message.warning('获取文章列表失败');
                    this.loading = false;
                }
            }).catch(err => {
                this.loading = false;
                window.$message.warning("获取文章列表出错");
            })
        },
        buildList(data) {
            if (data == null) data = []

            let num = 0
            let list = []
            for (let i in data) {
                let item = data[i]
                if (item.isBanner || item.isUp) continue
                list.push(item)
                num++
            }
            this.next = num == this.number
            if (data.length == 0) {
                window.$message.warning('已经到底啦');
                return false;
            }
            for (let i in list) {
                this.list.push(list[i])
            }
            if (this.next) this.page = this.page + 1;
        },
        open(id) {
            this.$router.push('/info/' + id)
        },
        openTag(id){
            this.$router.push('/list?tag='+id)
        },
        openGithub() {
            window.open('https://github.com/' + this.user.username)
        }
    },
    mounted() {
        this.init()
    },
};
</script>

<style scoped>
.welcome {
    font-size: 32px;
    font-weight: bold;
}

.readme {
    height: 117px;
}

.readme-info {
    height: 84px;
}

.readme-avatar {
    min-width: 40px;
    width: 40px;
    height: 40px;
}

.bind-tips {
    padding: 30px 0;
}

.tag-cloud {
    height: 172px;
}

.tag-item{
    cursor: pointer;
}

.article-box{
    min-height: calc(100vh - 431px);
}

.article-list {
    height: 200px;
}

.load-more {
    width: 40%;
}

.readme:deep(.social-btn) {
    width: 25%;
}
</style>
