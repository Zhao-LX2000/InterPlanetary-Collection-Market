<template>
    <div>
        
    </div>
</template>
<script lang="ts" setup>
    import { ElNotification } from 'element-plus' 
    // 3.服务器每次返回信息都会执行一次onmessage方法
    const open1 = (ms) => {
    ElNotification({
        title: '集群接入',
        dangerouslyUseHTMLString: true,
        message: ms,
        duration: 7000,
        type: 'success',
    })
    }

</script>

<style>
</style>