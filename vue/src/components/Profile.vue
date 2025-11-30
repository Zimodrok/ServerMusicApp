<template>
  <div class="fixed">
    <div
      class="bg-stone-100 dark:bg-stone-900 text-stone-800 dark:text-stone-300 font-sans relative md:fixed overflow-auto w-screen"
    >
      <div class="h-screen flex flex-col md:flex-row">
        <!-- Sidebar -->
        <div
          class="lg:w-custom-76 md:w-48 w-dvw m-2 mr-0 bg-white/80 dark:bg-stone-800/80 backdrop-blur-xl border-r border-stone-200 dark:border-stone-700 rounded-2xl scroll-smooth"
        >
          <div class="p-6 hidden md:block">
            <p
              class="text-xl font-semibold text-stone-900 mb-8 dark:text-white"
            >
              Music
            </p>
            <nav class="space-y-2 text-lg font-medium">
              <!-- <a
                href="#"
                class="items-center px-3 py-2 text-stone-700 rounded-lg flex space-x-3 dark:text-stone-300 hover:bg-stone-100 dark:hover:bg-stone-700"
              >
                <svg
                  class="w-8 h-6 fill-black dark:fill-white"
                  viewBox="0 0 38 28"
                  preserveAspectRatio="xMidYMid meet"
                >
                  <path
                    d="M0 10.2337C0 15.875 4.58576 20.4639 10.2337 20.4639C12.3037 20.4639 14.2165 19.8282 15.8074 18.7534L21.5753 24.5249C21.9374 24.8968 22.4376 25.0765 22.9506 25.0765C24.0575 25.0765 24.8515 24.2297 24.8515 23.1526C24.8515 22.6391 24.6717 22.1683 24.31 21.7999L18.5839 16.0506C19.766 14.4211 20.4706 12.4075 20.4706 10.2337C20.4706 4.58576 15.8817 0 10.2337 0C4.58576 0 0 4.58576 0 10.2337ZM2.60873 10.2337C2.60873 6.02948 6.02948 2.60873 10.2337 2.60873C14.4411 2.60873 17.8521 6.02948 17.8521 10.2337C17.8521 14.4313 14.4411 17.8521 10.2337 17.8521C6.02948 17.8521 2.60873 14.4313 2.60873 10.2337Z"
                    fill-opacity="0.85"
                  /></svg
                ><span>Search</span>
              </a> -->
              <a
                @click.prevent="$router.push('/library')"
                class="items-center px-3 py-2 text-stone-700 rounded-lg flex space-x-3 dark:text-stone-300 hover:bg-stone-100 dark:hover:bg-stone-700"
              >
                <svg
                  class="w-8 h-8 fill-black dark:fill-white"
                  viewBox="0 0 38 28"
                  preserveAspectRatio="xMidYMid meet"
                >
                  <path
                    d="M11.7188 25.095L18.949 25.095L18.949 17.1785C18.949 16.6036 18.5788 16.2366 18.007 16.2366L12.6705 16.2366C12.0988 16.2366 11.7188 16.6036 11.7188 17.1785ZM1.27864 13.8405C1.64614 13.8405 1.99276 13.6532 2.28835 13.3993L14.7454 2.94828C14.9291 2.79028 15.1355 2.71305 15.329 2.71305C15.5323 2.71305 15.7356 2.79028 15.9224 2.94828L28.3794 13.3993C28.675 13.6532 29.0119 13.8405 29.3891 13.8405C30.1734 13.8405 30.6678 13.2995 30.6678 12.6493C30.6678 12.2894 30.4983 11.9206 30.1672 11.641L17.1332 0.706149C16.5753 0.235233 15.9539 0 15.329 0C14.7139 0 14.0925 0.235233 13.5346 0.706149L0.500622 11.641C0.166425 11.9206 0 12.2894 0 12.6493C0 13.2995 0.487751 13.8405 1.27864 13.8405ZM23.3206 6.63102L26.8687 9.60566L26.8687 3.67326C26.8687 3.1384 26.5185 2.79795 25.9867 2.79795L24.2154 2.79795C23.6806 2.79795 23.3206 3.1384 23.3206 3.67326ZM6.8236 27.0148L23.8442 27.0148C25.7586 27.0148 26.8687 25.9145 26.8687 24.0666L26.8687 10.0322L24.4011 8.33102L24.4011 23.2772C24.4011 24.1068 23.9643 24.5503 23.1377 24.5503L7.52342 24.5503C6.70041 24.5503 6.26672 24.1068 6.26672 23.2772L6.26672 8.34434L3.79912 10.0322L3.79912 24.0666C3.79912 25.9211 4.90921 27.0148 6.8236 27.0148Z"
                    fill-opacity="0.85"
                  />
                </svg>
                <span>Home</span>
              </a>
              <!-- <a
                class="items-center px-3 py-2 text-stone-700 rounded-lg flex space-x-3 dark:text-stone-300 hover:bg-stone-100 dark:hover:bg-stone-700"
              >
                <svg
                  class="w-8 h-8 fill-black dark:fill-white"
                  viewBox="0 0 38 28"
                  preserveAspectRatio="xMidYMid meet"
                >
                  <path
                    d="M0 23.3816C0 25.8979 1.44204 27.3275 4.08837 27.3275L25.0773 27.3275C27.7743 27.3275 29.1897 25.912 29.1897 23.2479L29.1897 4.08939C29.1897 1.41859 27.7743 0.00976562 25.0773 0.00976562L9.81495 0.00976562C7.11796 0.00976562 5.70247 1.41859 5.70247 4.08939L5.70247 9.75161L3.02563 9.75161C1.12779 9.75161 0 10.8235 0 12.6454ZM2.4676 23.2108L2.4676 13.0728C2.4676 12.5468 2.80184 12.2161 3.33449 12.2161L5.70247 12.2161L5.70247 23.2108C5.70247 24.1758 5.01128 24.8599 4.08837 24.8599C3.15879 24.8599 2.4676 24.1528 2.4676 23.2108ZM7.78027 24.8599C8.03466 24.4257 8.17007 23.9116 8.17007 23.3025L8.17007 4.33921C8.17007 3.10376 8.79736 2.47737 10.0088 2.47737L24.8798 2.47737C26.0949 2.47737 26.7221 3.10376 26.7221 4.33921L26.7221 22.9981C26.7221 24.2335 26.0949 24.8599 24.8798 24.8599Z"
                    fill-opacity="0.85"
                  />
                  <path
                    d="M11.4163 7.87032L23.4923 7.87032C23.9512 7.87032 24.3063 7.50858 24.3063 7.0563C24.3063 6.60402 23.9575 6.25515 23.4923 6.25515L11.4163 6.25515C10.9445 6.25515 10.5957 6.60402 10.5957 7.0563C10.5957 7.50858 10.9476 7.87032 11.4163 7.87032ZM11.4163 12.2161L23.4923 12.2161C23.9575 12.2161 24.3063 11.8703 24.3063 11.4247C24.3063 10.9693 23.9512 10.6107 23.4923 10.6107L11.4163 10.6107C10.9476 10.6107 10.5957 10.9693 10.5957 11.4247C10.5957 11.8703 10.9445 12.2161 11.4163 12.2161ZM12.1319 20.9432L15.2362 20.9432C16.2012 20.9432 16.7955 20.3489 16.7955 19.3839L16.7955 16.5256C16.7955 15.5704 16.2012 14.9694 15.2362 14.9694L12.1319 14.9694C11.1638 14.9694 10.5628 15.5704 10.5628 16.5256L10.5628 19.3839C10.5628 20.3489 11.1638 20.9432 12.1319 20.9432ZM19.0346 16.5814L23.4856 16.5814C23.9446 16.5814 24.2934 16.2326 24.2934 15.79C24.2934 15.328 23.9446 14.9694 23.4856 14.9694L19.0346 14.9694C18.5627 14.9694 18.2139 15.328 18.2139 15.79C18.2139 16.2326 18.5658 16.5814 19.0346 16.5814ZM19.0346 20.9432L23.4856 20.9432C23.9446 20.9432 24.2934 20.5944 24.2934 20.1518C24.2934 19.6964 23.9446 19.3378 23.4856 19.3378L19.0346 19.3378C18.5658 19.3378 18.2139 19.6964 18.2139 20.1518C18.2139 20.5944 18.5627 20.9432 19.0346 20.9432Z"
                    fill-opacity="0.85"
                  /></svg
                ><span>New</span>
              </a> -->
              <!-- <a
                @click.prevent="$router.push('/profile')"
                class="items-center px-3 py-2 text-stone-700 rounded-lg flex space-x-3 dark:text-stone-300 hover:bg-stone-100 dark:hover:bg-stone-700"
              >
                <svg
                  class="w-8 h-8 fill-black dark:fill-white"
                  viewBox="0 0 38 28"
                  preserveAspectRatio="xMidYMid meet"
                >
                  <path
                    d="M12.7148 25.4395C19.7363 25.4395 25.4395 19.7461 25.4395 12.7246C25.4395 5.70312 19.7363 0 12.7148 0C5.69336 0 0 5.70312 0 12.7246C0 19.7461 5.69336 25.4395 12.7148 25.4395ZM12.7148 23.623C6.68945 23.623 1.81641 18.75 1.81641 12.7246C1.81641 6.69922 6.68945 1.82617 12.7148 1.82617C18.7402 1.82617 23.6133 6.69922 23.6133 12.7246C23.6133 18.75 18.7402 23.623 12.7148 23.623ZM21.2988 20.9668L21.2695 20.8496C20.5957 19.0332 17.2949 16.9629 12.7148 16.9629C8.1543 16.9629 4.85352 19.0234 4.16016 20.8301L4.13086 20.9668C6.38672 23.2227 10.0684 24.5703 12.7246 24.5703C15.3809 24.5703 19.0039 23.252 21.2988 20.9668ZM12.7148 14.8145C15.1367 14.834 17.0215 12.7637 17.0215 10.0684C17.0215 7.5293 15.1172 5.41992 12.7148 5.41992C10.3125 5.41992 8.39844 7.5293 8.4082 10.0684C8.41797 12.7637 10.293 14.7949 12.7148 14.8145Z"
                    fill-opacity="0.85"
                  />
                  d="M11.4163 7.87032L23.4923 7.87032C23.9512 7.87032 24.3063
                  7.50858 24.3063 7.0563C24.3063 6.60402 23.9575 6.25515 23.4923
                  6.25515L11.4163 6.25515C10.9445 6.25515 10.5957 6.60402
                  10.5957 7.0563C10.5957 7.50858 10.9476 7.87032 11.4163
                  7.87032ZM11.4163 12.2161L23.4923 12.2161C23.9575 12.2161
                  24.3063 11.8703 24.3063 11.4247C24.3063 10.9693 23.9512
                  10.6107 23.4923 10.6107L11.4163 10.6107C10.9476 10.6107
                  10.5957 10.9693 10.5957 11.4247C10.5957 11.8703 10.9445
                  12.2161 11.4163 12.2161ZM12.1319 20.9432L15.2362
                  20.9432C16.2012 20.9432 16.7955 20.3489 16.7955
                  19.3839L16.7955 16.5256C16.7955 15.5704 16.2012 14.9694
                  15.2362 14.9694L12.1319 14.9694C11.1638 14.9694 10.5628
                  15.5704 10.5628 16.5256L10.5628 19.3839C10.5628 20.3489
                  11.1638 20.9432 12.1319 20.9432ZM19.0346 16.5814L23.4856
                  16.5814C23.9446 16.5814 24.2934 16.2326 24.2934 15.79C24.2934
                  15.328 23.9446 14.9694 23.4856 14.9694L19.0346 14.9694C18.5627
                  14.9694 18.2139 15.328 18.2139 15.79C18.2139 16.2326 18.5658
                  16.5814 19.0346 16.5814ZM19.0346 20.9432L23.4856
                  20.9432C23.9446 20.9432 24.2934 20.5944 24.2934
                  20.1518C24.2934 19.6964 23.9446 19.3378 23.4856
                  19.3378L19.0346 19.3378C18.5658 19.3378 18.2139 19.6964
                  18.2139 20.1518C18.2139 20.5944 18.5627 20.9432 19.0346
                  20.9432Z" fill-opacity="0.85" /></svg
                ><span>Profile</span>
              </a> -->
            </nav>
          </div>
        </div>
        <div class="flex-1 flex flex-col relative h-screen">
          <div class="max-w-2xl mx-auto p-6">
            <div class="flex items-center justify-start mb-6 -translate-x-14">
              <button
                @click.prevent="$router.push('/library')"
                class="md:hidden w-12 h-12 rounded-full bg-stone-100 text-white flex items-center justify-center shadow-lg hover:bg-stone-300 transition-colors"
                aria-label="Profile"
              >
                <svg
                  viewBox="-4 -2 21 21"
                  class="w-6 h-6 fill-stone-700 dark:fill-white"
                >
                  <g transform="scale(0.85)">
                    <path
                      d="M0 10.6543C0 10.9277 0.107422 11.1621 0.3125 11.3672L10.3516 21.0254C10.5371 21.2012 10.7715 21.3086 11.0449 21.3086C11.5918 21.3086 12.0215 20.8887 12.0215 20.332C12.0215 20.0586 11.9141 19.8242 11.7383 19.6484L2.38281 10.6543L11.7383 1.65039C11.9141 1.47461 12.0215 1.23047 12.0215 0.966797C12.0215 0.410156 11.5918 0 11.0449 0C10.7715 0 10.5371 0.0976562 10.3516 0.283203L0.3125 9.94141C0.107422 10.1367 0 10.3711 0 10.6543Z"
                      fill-opacity="0.85"
                    />
                  </g>
                </svg>
              </button>

              <h1
                class="text-3xl ml-2 font-bold text-stone-800 dark:text-stone-100"
              >
                Profile
              </h1>
            </div>

            <form @submit.prevent="updateProfile" class="space-y-4">
              <div>
                <label class="block text-lg font-medium mb-1">Username</label>
                <input
                  type="text"
                  v-model="user.username"
                  disabled
                  class="input"
                />
              </div>

              <div class="grid grid-cols-2 gap-4">
                <div>
                  <label class="block text-lg font-medium mb-1"
                    >First Name</label
                  >
                  <input type="text" v-model="user.first_name" class="input" />
                </div>
                <div>
                  <label class="block text-lg font-medium mb-1"
                    >Last Name</label
                  >
                  <input type="text" v-model="user.last_name" class="input" />
                </div>
              </div>

              <div>
                <label class="block text-lg font-medium mb-1">Email</label>
                <input type="email" v-model="user.email" class="input" />
              </div>

              <div>
                <label class="block text-lg font-medium mb-1">Server IP</label>
                <input type="text" v-model="user.server_ip" class="input" />
              </div>

              <div>
                <label class="block text-lg font-medium mb-1"
                  >Library Path</label
                >
                <input type="text" v-model="user.library_path" class="input" />
              </div>

              <div class="flex justify-between mt-6">
                <button
                  type="submit"
                  class="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"
                >
                  Save Changes
                </button>
                <button
                  type="button"
                  @click="logout"
                  class="px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600"
                >
                  Log Out
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { getApiBase } from "../apiBase";

const user = ref({
  server_ip: "",
  server_port: "",
  library_path: "",
});

onMounted(async () => {
  const api = getApiBase();
  const res = await fetch(`${api}/profile`, {
    credentials: "include",
  });
  if (res.ok) {
    const data = await res.json();

    const [ip, port] = data.server?.split(":") || ["", ""];
    user.value = {
      ...data,
      server_ip: data.server, // full address:port for input
      server_port: port, // separate port if needed
      library_path: data.library_path, // correct path
    };

    console.log("Debug: user loaded", user.value);
  }
});

async function updateProfile() {
  // Combine server_ip and server_port for backend
  const payload = {
    ...user.value,
    server: `${user.value.server_ip}:${user.value.server_port}`,
    library_path: user.value.library_path,
  };

  const api = getApiBase();
  const res = await fetch(`${api}/profile`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    credentials: "include",
    body: JSON.stringify(payload),
  });

  if (res.ok) alert("Profile updated successfully!");
}

async function logout() {
  const userId = user.value.id;
  const api = getApiBase();
  await fetch(`${api}/logout`, {
    method: "POST",
    credentials: "include",
  });
  if (userId) localStorage.removeItem(`uploadGuideShown_${userId}`);
  window.location.href = "/";
}
</script>
<style scoped>
.input {
  @apply w-full border rounded-lg px-3 py-2 dark:bg-stone-800 dark:border-stone-700 dark:text-stone-100;
}
</style>
