<template>
    <n-drawer v-model:show="show" :width="370" :close-on-esc="false" content-class="file-box" placement="right">
        <n-drawer-content title="文件管理">
            <div v-if="loading" class="loading-block"><n-spin /></div>
            <n-scrollbar style="height: calc(100vh - 55px)">
                <div class="flex">
                    <div class="img-list" v-if="list1.length > 0" style="margin-right: -1px;">
                        <div v-for="item in list1" class="img-item pa-10 border-bottom border-right">
                            <n-button class="img-rm" size="small" strong secondary type="error" @click.stop="remove(item.name)">
                                <template #icon>
                                    <n-icon>
                                        <Delete16Regular />
                                    </n-icon>
                                </template>
                            </n-button>
                            <n-image class="img-item" width="160" :src="'/res/' + item.name" />
                            <div class="flex align-center justify-between">
                                <n-time class="text-small text-gray" :time="item.time" format="yyyy-MM-dd hh:mm" />
                                <div class="text-small text-gray">{{ getSize(item.size) }}</div>
                            </div>
                        </div>
                    </div>
                    <div class="img-list" v-if="list2.length > 0">
                        <div v-for="item in list2" class="img-item pa-10 border-bottom border-left">
                            <n-button class="img-rm" size="small" strong secondary type="error" @click.stop="remove(item.name)">
                                <template #icon>
                                    <n-icon>
                                        <Delete16Regular />
                                    </n-icon>
                                </template>
                            </n-button>
                            <n-image class="img-item" width="160" :src="'/res/' + item.name" />
                            <div class="flex align-center justify-between">
                                <n-time class="text-small text-gray" :time="item.time" format="yyyy-MM-dd hh:mm" />
                                <div class="text-small text-gray">{{ getSize(item.size) }}</div>
                            </div>
                        </div>
                    </div>
                </div>
            </n-scrollbar>
        </n-drawer-content>
    </n-drawer>
</template>
<script>
import { common } from '../plugins/api'
import { Delete16Regular } from '@vicons/fluent'

export default {
    name: "FilesEdit",
    components: { Delete16Regular },
    data: () => ({
        show: false,
        loading: false,
        list1: [],
        list2: [],
    }),
    methods: {
        open() {
            this.loading = true
            this.init()
            this.show = true
        },
        init() {
            this.list1 = []
            this.list2 = []
            common.getFiles().then(res => {
                if (res.state) {
                    let jump = true
                    for (let i in res.data) {
                        if (jump) this.list1.push(res.data[i])
                        else this.list2.push(res.data[i])
                        jump = !jump
                    }
                    console.log(this.list1,this.list2)
                }
                this.loading = false
            }).catch(err => {
                window.$message.warning("发生意料之外的错误");
                this.loading = false
            })
        },
        remove(name) {
            window.$dialog.warning({
                title: '操作确认',
                content: '文件删除不可逆, 确认要删除此文件吗?',
                positiveText: '确认删除',
                negativeText: '取消',
                onPositiveClick: () => {
                    common.removeFile(name).then(res => {
                        if (res.state) {
                            this.init();
                            window.$message.success('文件删除成功');
                        }
                        else window.$message.warning('文件删除失败');
                    }).catch(err => {
                        window.$message.warning("发生意料之外的错误");
                    })
                }
            })
        },
        getSize(size) {
            if (size < 1024) return size + ' B';
            else if (size < (1024 * 1024)) return parseInt(size / 1024) + ' KB';
            else if (size < (1024 * 1024 * 1024)) return parseInt(size / 1024 / 1024) + ' MB';
        }
    }
};
</script>
<style scoped>
.img-list {
    width: 185px;
}

.img-item{
    position: relative;
}

.img-rm{
    position: absolute;
    z-index: 99;
    right: 5px;
    top: 5px;
}
</style>