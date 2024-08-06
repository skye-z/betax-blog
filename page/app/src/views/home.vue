<template>
    <div class="app-content no-select">
        <div class="flex mb-10">
            <div class="card carousel mr-10">
                <n-carousel direction="vertical" dot-placement="right" show-arrow style="border-radius: 8px;">
                    <img class="carousel-item"
                        src="https://naive-ui.oss-cn-beijing.aliyuncs.com/carousel-img/carousel1.jpeg">
                    <div class="card carousel-item pa-10">文章</div>
                </n-carousel>
            </div>
            <div class="full-width">
                <div class="card readme mb-10 full-width">
                    <div class="flex pa-10 border-bottom">
                        <n-avatar class="mr-10" size="large"
                            src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg" />
                        <div>
                            <div>Skye Zhang</div>
                            <div class="text-gray text-small">@skye-z</div>
                        </div>
                    </div>
                    <div>
                        <n-button class="social-btn" quaternary>
                            <template #icon>
                                <n-icon>
                                    <Github />
                                </n-icon>
                            </template>
                        </n-button>
                        <n-button class="social-btn" quaternary disabled>
                            <template #icon>
                                <n-icon>
                                    <QqOutlined />
                                </n-icon>
                            </template>
                        </n-button>
                        <n-button class="social-btn" quaternary disabled>
                            <template #icon>
                                <n-icon>
                                    <Linkedin />
                                </n-icon>
                            </template>
                        </n-button>
                        <n-button class="social-btn" quaternary disabled>
                            <template #icon>
                                <n-icon>
                                    <Discord />
                                </n-icon>
                            </template>
                        </n-button>
                    </div>
                </div>
                <div class="card tag-cloud pa-10 full-width">标签云</div>
            </div>
        </div>
        <div class="article-box">
            <div class="card article-list mb-10 pa-10">文章</div>
            <div class="card article-list mb-10 pa-10">文章</div>
            <div class="card article-list mb-10 pa-10">文章</div>
            <div class="card article-list mb-10 pa-10">文章</div>
            <div class="card article-list mb-10 pa-10">文章</div>
            <div class="text-center">
                <n-button class="load-more" strong secondary>加载更多</n-button>
            </div>
        </div>
    </div>
</template>

<script>
import { Github, Linkedin, Discord } from '@vicons/fa'
import { QqOutlined } from '@vicons/antd'

export default {
    name: "Home",
    components: { Github, Linkedin, QqOutlined, Discord },
    data: () => ({

    }),
    methods: {
        init() {

        }
    },
    mounted() {
        this.init()
    },
};
</script>

<style scoped>
.carousel {
    width: 580px;
    height: 300px;
}

.carousel-item {
    width: 580px;
    height: 300px;
}

.readme {
    height: 97px;
}

.tag-cloud {
    height: 192px;
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
