<template>
  <el-header class="flex-inline w-full align-items">
    <el-row justify="center" align="middle">
      <el-col :span="8">
        <el-menu mode="horizontal" class="font-bold" active-text-color="#06b6d4">
          <el-menu-item index="1">
            <router-link to="/">Home</router-link>
          </el-menu-item>
          <el-menu-item index="2">Documents</el-menu-item>
        </el-menu>
      </el-col>
      <el-col :span="16" class="text-right content-center">
        <el-button type="text" @click="notImplement('Search')">
          <search theme="outline" size="18" fill="#333"/>
        </el-button>
        <el-button type="text">
          <a :href="githuburl" target="_blank">
          <github-one theme="outline" size="18" :fill="['#333']"/>
          </a>
        </el-button>
        <el-dropdown placement="bottom" class="mx-1rem">
          <me theme="two-tone" size="18" :fill="['#333' ,'#50e3c2']"/>
          <template #dropdown>
            <div class="flex items-center content-center text-center mt-1rem mb-0.5rem mx-4 font-bold text-l">
               <sun theme="two-tone" size="24" :fill="['#333' ,'#f8e71c']"/>
               <span class="ml-2">Hi {{ username }}</span>
            </div>
            <el-dropdown-menu>
              <el-dropdown-item :icon="Info" @click="notImplement('UserInfo')">
                UserInfo
              </el-dropdown-item>
            </el-dropdown-menu>
            <el-dropdown-menu>
              <el-dropdown-item :icon="SettingOne" @click="notImplement('Setting')">
                Setting
              </el-dropdown-item>
            </el-dropdown-menu>

            <el-dropdown-menu>
              <el-dropdown-item :icon="Logout" @click="logout">
                Logout
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </el-col>
    </el-row>
  </el-header>
</template>

<script setup>
import { Info, SettingOne, Logout, Sun, Search, GithubOne, Success, Me } from '@icon-park/vue-next';
import { getUser } from '@/utils';
import request from '@/axios'
import { ElMessage, ElNotification } from "element-plus";
import { useRouter } from 'vue-router';

const username = getUser();
const githuburl = 'https://github.com/qingwave/weave';
const router = useRouter();

function logout() {
  request.delete("/api/auth/token").then(() => {
    console.log("logout success")
    ElNotification.success({
        title: 'Logout Success',
        message: 'Bye~ ' + username,
        showClose: false,
    })
  }).catch((error) => {
    console.log(error)
  })
  router.push('/login');
};

function notImplement(name) {
  ElMessage({
    message: name + ' Coming Soon',
    type: 'warning',
    duration: 1000,
  })
}

</script>

<style scoped>
.el-header {
  background-color: white;
  margin-bottom: 1px;
  box-shadow:1px 1px 1px #dcdbdb;
}
</style>