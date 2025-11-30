<template>
  <div class="w-full">
    <div
      class="w-full bg-white text-neutral-800 dark:bg-neutral-950 relative isolate antialiased dark:text-neutral-100 min-h-screen"
    >
      <div
        class="w-full pointer-events-none absolute inset-0 -z-10 overflow-hidden"
      >
        <div
          class="h-[60vh] w-[60vh] rounded-full bg-gradient-to-br absolute -top-32 -left-32 from-indigo-200 via-lime-200 to-purple-300 dark:from-orange-600 dark:via-amber-500 dark:to-rose-400 opacity-20 blur-2xl dark:opacity-80"
        ></div>
        <div
          class="h-[40vh] w-[50vh] rounded-full bg-gradient-to-tr absolute -bottom-20 right-10 from-fuchsia-300 via-orange-300 to-rose-200 dark:from-orange-600 dark:via-amber-500 dark:to-rose-400 opacity-40 blur-3xl dark:opacity-80"
        ></div>
        <div
          class="h-[35vh] w-[45vh] rounded-full bg-gradient-to-b absolute top-28 left-1/4 from-orange-300 via-amber-200 to-rose-100 opacity-60 blur-3xl dark:from-orange-600 dark:via-amber-500 dark:to-rose-400 dark:opacity-64"
        ></div>
      </div>
      <header
        class="mx-auto items-center justify-between px-8 pt-10 text-lg relative z-10 flex max-w-7xl"
      >
        <div class="items-center flex gap-3">
          <div
            class="w-8 h-8 rounded-lg bg-neutral-900 dark:bg-neutral-100 items-center justify-center flex"
          >
            <svg
              class="w-5 h-5 text-white dark:text-neutral-900"
              fill="currentColor"
              viewBox="0 0 24 24"
              id="Windframe_ugUWOHBMk"
            >
              <path
                d="M12 3v10.55c-.59-.34-1.27-.55-2-.55-2.21 0-4 1.79-4 4s1.79 4 4 4 4-1.79 4-4V7h4V3h-6z"
              ></path>
            </svg>
          </div>
          <p class="font-semibold tracking-wide text-lg">FlacLibrary</p>
        </div>
        <nav class="lg:flex items-center hidden gap-8">
          <a
            href="/Features.html"
            class="hover:text-neutral-900 dark:hover:text-neutral-100 dark:text-neutral-300 transition-colors"
            >Features</a
          >
          <a
            href="https://github.com/Zimodrok/InformNetw-public"
            class="hover:text-neutral-900 dark:hover:text-neutral-100 dark:text-neutral-300 transition-colors"
            >Project Page</a
          >
          <button
            v-if="showInitServer"
            type="button"
            class="z-[120]"
            @click="onInitServer"
            :class="[
              'inline-flex items-center justify-center rounded-md border border-neutral-500 px-2 ' +
                'hover:text-neutral-900 dark:hover:text-neutral-100 dark:text-neutral-300 ' +
                'transition-colors hover:bg-neutral-100 dark:hover:bg-neutral-800 ' +
                'disabled:opacity-60 disabled:cursor-not-allowed',
              initServerHighlight
                ? 'dock-bounce bg-neutral-900 text-white shadow-xl z-50'
                : '',
            ]"
            :disabled="initServerLoading"
          >
            <span v-if="!initServerLoading">Init Server</span>
            <span v-else>Initializing…</span>
          </button>

          <p v-if="initServerError" class="mt-2 text-sm text-red-400">
            {{ initServerError }}
          </p>
          <input
            ref="folderInput"
            type="file"
            webkitdirectory
            directory
            multiple
            class="hidden"
            @change="onFolderPicked"
          />
        </nav>
        <div class="lg:flex items-center hidden gap-4">
          <a
            href="#signin"
            @click.prevent="showLoginPopup = true"
            class="text-neutral-900 font-medium dark:text-neutral-100 hover:underline"
            >Sign in</a
          >
        </div>
      </header>
      <main class="mx-auto px-8 relative z-20 max-w-4xl">
        <div class="lg:grid-cols-2 lg:py-24 items-center py-16 grid gap-16">
          <div class="space-y-8">
            <div class="space-y-6">
              <div
                class="items-center px-3 py-1 rounded-full bg-white/70 text-sm dark:bg-white/10 inline-flex gap-2 border border-neutral-300/70 dark:border-white/20"
              >
                <span class="w-2 h-2 bg-green-500 rounded-full"></span>
                <span class="text-neutral-700 dark:text-neutral-300"
                  >High-Quality Audio</span
                >
              </div>
              <p class="text-4xl font-medium leading-tight lg:text-5xl">
                Join FLAC Music Player
              </p>
              <p
                class="text-lg text-neutral-700/80 dark:text-neutral-300/80 max-w-lg"
              >
                Create your account to start managing your local music library
                with lossless audio quality. Organize, play, and discover your
                music collection like never before.
              </p>
            </div>
          </div>
          <div class="relative">
            <div
              class="bg-white/80 dark:bg-neutral-900/80 rounded-2xl shadow-2xl backdrop-blur-lg border border-neutral-200/50 dark:border-neutral-700/50 p-8"
            >
              <div class="space-y-6" v-if="showRegisterForm">
                <div class="text-center space-y-2">
                  <p class="text-2xl font-semibold">Create Account</p>
                  <p class="text-neutral-600 dark:text-neutral-400">
                    Start your music journey today
                  </p>
                </div>
                <form class="space-y-3" @submit.prevent="submitRegister">
                  <div>
                    <div class="sm:grid-cols-2 grid grid-cols-1 gap-4">
                      <div class="space-y-2">
                        <label
                          for="firstName"
                          class="text-md font-medium text-neutral-700 dark:text-neutral-300"
                          >First Name</label
                        >
                        <input
                          v-model="registerForm.firstName"
                          name="firstName"
                          type="text"
                          placeholder="John"
                          class="border border-neutral-300/70 dark:border-neutral-600 focus:border-neutral-500 dark:focus:border-neutral-400 focus:outline-none transition-colors w-full px-4 py-3 rounded-lg bg-white/50 dark:bg-neutral-800/50"
                          id="firstName"
                          maxlength="40"
                          title="Only letters and legit characters are allowed."
                        />
                      </div>
                      <div class="space-y-2">
                        <label
                          for="lastName"
                          class="text-md font-medium text-neutral-700 dark:text-neutral-300"
                          >Last Name</label
                        >
                        <input
                          v-model="registerForm.lastName"
                          name="lastName"
                          type="text"
                          placeholder="Doe"
                          class="border border-neutral-300/70 dark:border-neutral-600 focus:border-neutral-500 dark:focus:border-neutral-400 focus:outline-none transition-colors w-full px-4 py-3 rounded-lg bg-white/50 dark:bg-neutral-800/50"
                          id="lastName"
                          maxlength="40"
                          title="Only letters and legit characters are allowed."
                        />
                      </div>
                    </div>
                    <p v-if="errors.name" class="text-red-500 text-md mt-1">
                      {{ errors.name }}
                    </p>
                    <p
                      class="text-md text-neutral-600 mt-2 w-full dark:text-neutral-400 text-center"
                    >
                      Optional
                    </p>
                  </div>
                  <div class="space-y-2">
                    <label
                      for="email"
                      class="text-md font-medium text-neutral-700 dark:text-neutral-300"
                      >Email Address</label
                    >
                    <input
                      v-model="registerForm.email"
                      name="email"
                      required=""
                      type="email"
                      placeholder="john@example.com"
                      class="border border-neutral-300/70 dark:border-neutral-600 focus:border-neutral-500 dark:focus:border-neutral-400 focus:outline-none transition-colors w-full px-4 py-3 rounded-lg bg-white/50 dark:bg-neutral-800/50"
                      id="email"
                      maxlength="80"
                      pattern="^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[A-Za-z]{2,}$"
                      title="Enter a valid email address (letters, digits, ., _, %, +, -, and @)"
                    />
                  </div>
                  <p v-if="errors.email" class="text-red-500 text-md mt-1">
                    {{ errors.email }}
                  </p>

                  <div class="space-y-2">
                    <label
                      for="username"
                      class="text-md font-medium text-neutral-700 dark:text-neutral-300"
                      >Username</label
                    >
                    <input
                      v-model="registerForm.username"
                      name="username"
                      required=""
                      type="text"
                      placeholder="audiophile123"
                      class="border border-neutral-300/70 dark:border-neutral-600 focus:border-neutral-500 dark:focus:border-neutral-400 focus:outline-none transition-colors w-full px-4 py-3 rounded-lg bg-white/50 dark:bg-neutral-800/50"
                      id="username"
                      maxlength="40"
                      :class="{ 'border-red-500': errors.username }"
                    />
                    <p v-if="errors.username" class="text-red-500 text-md mt-1">
                      {{ errors.username }}
                    </p>
                  </div>
                  <div class="space-y-2">
                    <label
                      for="password"
                      class="text-md font-medium text-neutral-700 dark:text-neutral-300"
                      >Password</label
                    >
                    <input
                      v-model="registerForm.password"
                      name="password"
                      required
                      type="password"
                      placeholder="••••••••"
                      id="password"
                      class="border border-neutral-300/70 dark:border-neutral-600 focus:border-neutral-500 dark:focus:border-neutral-400 focus:outline-none transition-colors w-full px-4 py-3 rounded-lg bg-white/50 dark:bg-neutral-800/50"
                      :class="{ 'border-red-500': errors.password }"
                    />
                    <p v-if="errors.password" class="text-red-500 text-md mt-1">
                      {{ errors.password }}
                    </p>
                  </div>
                  <div class="space-y-2">
                    <label
                      for="confirmPassword"
                      class="text-md font-medium text-neutral-700 dark:text-neutral-300"
                      >Confirm Password</label
                    >
                    <input
                      v-model="registerForm.confirmPassword"
                      name="confirmPassword"
                      required
                      type="password"
                      placeholder="••••••••"
                      id="confirmPassword"
                      class="border border-neutral-300/70 dark:border-neutral-600 focus:border-neutral-500 dark:focus:border-neutral-400 focus:outline-none transition-colors w-full px-4 py-3 rounded-lg bg-white/50 dark:bg-neutral-800/50"
                      :class="{ 'border-red-500': errors.confirmPassword }"
                    />
                    <p
                      v-if="errors.confirmPassword"
                      class="text-red-500 text-md mt-1"
                    >
                      {{ errors.confirmPassword }}
                    </p>
                  </div>
                  <div class="space-y-2">
                    <label
                      for="libraryPath"
                      class="text-md font-medium text-neutral-700 dark:text-neutral-300"
                      >Music Library Path
                    </label>
                    <input
                      v-model="registerForm.libraryPath"
                      name="libraryPath"
                      type="text"
                      placeholder="/Users/john/Music"
                      class="border border-neutral-300/70 dark:border-neutral-600 focus:border-neutral-500 dark:focus:border-neutral-400 focus:outline-none transition-colors w-full px-4 py-3 rounded-lg bg-white/50 dark:bg-neutral-800/50"
                      id="libraryPath"
                      pattern="^[A-Za-z0-9_\/\\.\-\s]{1,200}$"
                      maxlength="100"
                      title="Letters, digits, slashes (/ or \\), dots (.), underscores (_), hyphens (-), and spaces only."
                    />
                    <p class="text-md text-neutral-600 dark:text-neutral-400">
                      Optional. You can set this up later in settings
                    </p>
                  </div>
                  <button
                    type="submit"
                    class="border border-transparent transition-colors hover:bg-neutral-700 dark:hover:bg-neutral-600 w-full items-center justify-center rounded-lg bg-neutral-900 px-8 py-3 font-medium text-neutral-100 dark:bg-neutral-700"
                  >
                    Create Account
                  </button>
                  <div class="relative">
                    <div class="items-center absolute inset-0 flex">
                      <div
                        class="w-full border-t border-neutral-300 dark:border-neutral-700"
                      ></div>
                    </div>
                  </div>
                  <!-- <div class="justify-center text-sm relative flex"> <span
                        class="px-2 bg-white dark:bg-neutral-900 text-neutral-500 dark:text-neutral-400">Or continue with</span>
                    </div>
                  </div>
                  <div class="grid grid-cols-2 gap-3"> <button type="button"
                      class="inline-flex border border-neutral-300 dark:border-neutral-700 dark:text-neutral-300 hover:bg-neutral-50 dark:hover:bg-neutral-700 transition-colors items-center justify-center px-4 py-2 rounded-lg bg-white dark:bg-neutral-800 text-neutral-700">
                      <svg class="w-5 h-5 mr-2" viewBox="0 0 24 24" id="Windframe_RaaetPnKJ">
                        <path fill="currentColor"
                          d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z">
                        </path>
                        <path fill="currentColor"
                          d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z">
                        </path>
                        <path fill="currentColor"
                          d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z">
                        </path>
                        <path fill="currentColor"
                          d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z">
                        </path>
                      </svg> Google </button> <button type="button"
                      class="inline-flex border border-neutral-300 dark:border-neutral-700 dark:text-neutral-300 hover:bg-neutral-50 dark:hover:bg-neutral-700 transition-colors items-center justify-center px-4 py-2 rounded-lg bg-white dark:bg-neutral-800 text-neutral-700">
                      <svg class="w-5 h-5 mr-2" fill="currentColor" viewBox="0 0 24 24" id="Windframe_k8jgpIQW0">
                        <path
                          d="M12.017 0C5.396 0 .029 5.367.029 11.987c0 5.079 3.158 9.417 7.618 11.174-.105-.949-.199-2.403.041-3.439.219-.937 1.406-5.957 1.406-5.957s-.359-.72-.359-1.781c0-1.663.967-2.911 2.168-2.911 1.024 0 1.518.769 1.518 1.688 0 1.029-.653 2.567-.992 3.992-.285 1.193.6 2.165 1.775 2.165 2.128 0 3.768-2.245 3.768-5.487 0-2.861-2.063-4.869-5.008-4.869-3.41 0-5.409 2.562-5.409 5.199 0 1.033.394 2.143.889 2.741.099.12.112.225.085.347-.09.375-.292 1.199-.334 1.363-.053.225-.172.271-.402.165-1.495-.69-2.433-2.878-2.433-4.646 0-3.776 2.748-7.252 7.92-7.252 4.158 0 7.392 2.967 7.392 6.923 0 4.135-2.607 7.462-6.233 7.462-1.214 0-2.357-.629-2.758-1.378l-.749 2.848c-.269 1.045-1.004 2.352-1.498 3.146 1.123.345 2.306.535 3.55.535 6.624 0 11.99-5.367 11.99-11.986C24.007 5.367 18.641.001 12.017.001z">
                        </path>
                      </svg> Apple </button> </div> -->
                </form>
                <div class="mt-8 text-center">
                  <p class="text-md text-neutral-600 dark:text-neutral-400">
                    Already have an account?
                    <a
                      href="#signin"
                      @click.prevent="showLoginPopup = true"
                      class="text-neutral-900 font-bold dark:text-neutral-100 hover:underline"
                      >Sign in</a
                    >
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>
      <div
        v-if="showLoginPopup"
        class="fixed inset-0 bg-black/50 flex items-center justify-center z-40"
      >
        <div
          class="bg-white dark:bg-neutral-900 rounded-xl shadow-xl p-6 w-full max-w-md relative"
        >
          <button
            @click="closeLoginPopup"
            class="absolute top-3 right-3 text-neutral-500 hover:text-neutral-800"
          >
            ✕
          </button>
          <p class="text-xl font-semibold mb-4">Sign In</p>
          <form class="space-y-4" @submit.prevent="submitLogin">
            <input
              v-model="loginForm.username"
              type="text"
              placeholder="Username"
              class="input-field bg-white dark:bg-neutral-700 dark:text-neutral-100 text-neutral-900"
            />
            <input
              v-model="loginForm.password"
              type="password"
              placeholder="Password"
              class="input-field bg-white dark:bg-neutral-700 dark:text-neutral-100 text-neutral-900"
            />
            <label
              class="flex items-center space-x-2 text-sm text-neutral-600 dark:text-neutral-300"
            >
              <input
                type="checkbox"
                v-model="loginForm.rememberMe"
                class="w-4 h-4"
              />
              <span>Remember me (30 days)</span>
            </label>
            <button
              type="submit"
              class="w-full rounded-lg bg-neutral-900 dark:bg-neutral-700 px-6 py-3 font-medium text-neutral-100 hover:bg-neutral-700 dark:hover:bg-neutral-600"
            >
              Sign In
            </button>
          </form>
          <button
            type="button"
            @click="loginAsGuest"
            class="w-full rounded-lg px-6 py-4 font-medium text-neutral-800 hover:text-black dark:text-neutral-100 font-semibold dark:hover:text-white"
          >
            Continue as Guest
          </button>
        </div>
      </div>
      <div
        v-if="sftpModalShown"
        class="fixed inset-0 z-40 flex items-center justify-center bg-black/50"
      >
        <div
          class="w-full max-w-lg mx-4 bg-white dark:bg-neutral-800 rounded-2xl shadow-lg p-6"
        >
          <div class="flex items-center justify-between mb-4">
            <h3
              class="text-lg font-semibold text-neutral-900 dark:text-neutral-100"
            >
              Connect your SFTP server
            </h3>
            <button
              @click="onSftpModalCancel"
              class="text-neutral-500 hover:text-neutral-700"
            >
              ✕
            </button>
          </div>

          <form @submit.prevent="submitSftpCreds" class="space-y-4">
            <div>
              <label
                class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1"
                >Host (host or host:port)</label
              >
              <input
                v-model="hostInput"
                type="text"
                placeholder="localhost:22"
                class="input-field"
                :class="
                  sftpExists
                    ? 'bg-neutral-200 dark:bg-neutral-800 dark:text-neutral-500 text-neutral-500'
                    : 'bg-white dark:bg-neutral-700 dark:text-neutral-100 text-neutral-900'
                "
              />
              <p v-if="hostError" class="text-sm text-red-500 mt-1">
                {{ hostError }}
              </p>
            </div>

            <div class="grid grid-cols-2 gap-4">
              <div>
                <label
                  class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1"
                  >SFTP Username</label
                >
                <input
                  v-model="sftpUser"
                  type="text"
                  class="input-field"
                  :class="
                    sftpExists
                      ? 'bg-neutral-200 dark:bg-neutral-800 dark:text-neutral-500 text-neutral-500'
                      : 'bg-white dark:bg-neutral-700 dark:text-neutral-100 text-neutral-900'
                  "
                  :placeholder="loginForm.username"
                />
              </div>
              <div>
                <label
                  class="block text-sm font-medium text-neutral-700 dark:text-neutral-300 mb-1"
                  >SFTP Password</label
                >
                <input
                  v-model="sftpPassword"
                  type="password"
                  class="input-field bg-white dark:bg-neutral-700 dark:text-neutral-100 text-neutral-900"
                />
              </div>
            </div>

            <div>
              <label
                class="block text-sm font-medium text-neutral-700 ne dark:text-neutral-300 mb-1"
                >Library path</label
              >
              <input
                v-model="libraryPath"
                type="text"
                class="input-field"
                :class="
                  sftpExists
                    ? 'bg-neutral-200 dark:bg-neutral-800 dark:text-neutral-500 text-neutral-500'
                    : 'bg-white dark:bg-neutral-700 dark:text-neutral-100 text-neutral-900'
                "
              />
            </div>

            <div
              v-if="serverInfo"
              class="text-sm text-neutral-600 dark:text-neutral-300"
            >
              Detected host:
              <span class="font-medium">{{ serverInfo.host }}</span> port:
              <span class="font-medium">{{ serverInfo.port }}</span>
            </div>

            <div class="flex items-center justify-end space-x-2 mt-2">
              <button
                type="button"
                @click="onSftpModalCancel"
                class="px-4 py-2 rounded-lg border border-neutral-200 dark:border-neutral-700"
              >
                Cancel
              </button>
              <button
                type="submit"
                :disabled="loading"
                class="px-4 py-2 bg-blue-600 text-white rounded-lg"
              >
                <span v-if="!loading">Connect</span>
                <span v-else>Connecting…</span>
              </button>
            </div>

            <p
              v-if="error"
              class="text-lg capitalize font-bold text-white bg-red-500 rounded-lg p-2 absolute"
            >
              {{ error }}
            </p>
            <p
              v-if="success"
              class="text-lg capitalize font-bold text-green-500 mt-2"
            >
              {{ success }}
            </p>
          </form>
        </div>
      </div>
      <div class="absolute top-20 right-20 z-40">
        <transition name="slide-right" appear mode="out-in">
          <div
            v-if="loginError"
            class="text-white p-3 rounded-xl text-xl bg-red-600 font-semibold bg-opacity-75 text-sm mb-2"
          >
            {{ loginError }}
          </div>
        </transition>
        <transition name="slide-right" appear mode="out-in">
          <div
            v-if="positiveRespone"
            class="text-white p-3 rounded-xl text-xl bg-neutral-600 font-semibold bg-opacity-75 text-sm mb-2"
          >
            {{ positiveRespone }}
          </div>
        </transition>
      </div>
    </div>
  </div>
</template>

<script>
import { getApiBase, getPortsConfig } from "../apiBase";
export default {
  name: "Auth",
  data() {
    const cfg = getPortsConfig() || {};
    const defaultPort = cfg.sftp_port || 9824;
    return {
      showRegisterForm: true,
      showLoginPopup: false,
      loginError: "",
      positiveRespone: "",
      sftpModalShown: false,
      sftpPromiseResolve: null,
      sftpPromiseReject: null,
      showInitServer: false,
      initServerBounce: false,
      initServerLoading: false,
      initServerError: "",
      sftpExists: false,
      defaultSftpPort: defaultPort,
      hostInput: `localhost:${defaultPort}`,
      sftpUser: "FlacPlayerUser",
      sftpPassword: "",
      localSftpFolder: "",
      localSftpFolderPath: "",
      waitingForFolderPick: false,
      loading: false,
      error: "",
      success: "",
      hostError: "",

      registerForm: {
        firstName: "",
        lastName: "",
        email: "",
        username: "",
        password: "",
        confirmPassword: "",
      },
      errors: {
        password: "",
        confirmPassword: "",
        username: "",
        name: "",
        email: "",
      },
      loginForm: { username: "", password: "", rememberMe: true },
    };
  },

  watch: {
    "registerForm.firstName"(val) {
      const regex = /^[\p{L}'’\-–]+$/u;
      this.errors.name =
        val && !regex.test(val) ? "Only letters are allowed" : "";
    },
    "registerForm.lastName"(val) {
      const regex = /^[\p{L}'’\-–]+$/u;
      this.errors.name =
        val && !regex.test(val) ? "Only letters are allowed" : "";
    },
    "registerForm.email"(val) {
      const regex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
      this.errors.email = val && !regex.test(val) ? "Email is not valid" : "";
    },
    "registerForm.username"(val) {
      const regex = /^[a-zA-Z0-9_.]{3,30}$/;
      this.errors.username =
        val && !regex.test(val) ? "3–30 chars, letters, digits, _ or ." : "";
    },
    "registerForm.password"(val) {
      const regex = /^[A-Za-z0-9!@#$%^&*()_\-=\\[\]{}|;:,.<>?/`~]{4,100}$/;
      this.errors.password =
        val && !regex.test(val) ? "Invalid password format" : "";
    },
    "registerForm.confirmPassword"(val) {
      this.errors.confirmPassword =
        val && val !== this.registerForm.password
          ? "Passwords do not match"
          : "";
    },
  },

  methods: {
    parseHost(input) {
      const re = /^([\w\.\-]+)(?::(\d{1,5}))?$/;
      const m = input.trim().match(re);
      if (!m) return null;
      const host = m[1];
      const port = m[2] ? parseInt(m[2], 10) : 22;
      if (port < 1 || port > 65535) return null;
      return { host, port };
    },
    bumpInitServerButton() {
      this.initServerHighlight = true;
      setTimeout(() => {
        this.initServerHighlight = false;
      }, 3000);
    },
    openFolderPicker() {
      this.$refs.folderInput?.click();
    },
    onFolderPicked(event) {
      const files = event.target.files;
      if (!files || !files.length) return;

      const first = files[0];
      const rel = first.webkitRelativePath || "";
      const rootName = rel.split("/")[0] || first.name;

      console.log("Picked folder name:", rootName);

      this.localSftpFolderName = rootName;

      if (this.waitingForFolderPick) {
        this.onInitServer();
      }
    },
    async submitSftpCreds() {
      const api = getApiBase();
      const parsed = this.parseHost(this.hostInput);
      if (!parsed.host || !parsed.port) {
        this.error = "Invalid host (use host:port)";
        return;
      }

      if (!this.sftpUser || !this.sftpPassword) {
        this.error = "Username and password required";
        return;
      }

      this.loading = true;
      this.showInitServer = false;

      try {
        const body = {
          host: parsed.host,
          port: Number(parsed.port),
          username: this.sftpUser,
          password: this.sftpPassword,
          path: this.libraryPath || `${this.loginForm.username}/library`,
        };
        console.log("submitSftpCreds");
        const res = await fetch(`${api}/sftp/creds`, {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          credentials: "include",
          body: JSON.stringify(body),
        });

        const data = await res.json().catch(() => ({}));
        if (!res.ok) {
          this.error = data.details || data.error || "SFTP connection failed";

          this.showInitServer = true;
          this.initServerHighlight = true;
          setTimeout(() => {
            this.initServerHighlight = false;
          }, 2000);

          setTimeout(() => (this.error = ""), 3000);
          this.loading = false;
          return;
        }

        this.success = "SFTP saved and reachable";
        setTimeout(() => {
          this.loading = false;
          this.sftpModalShown = false;
          this.onSftpModalConnected(body);
        }, 800);
      } catch (e) {
        console.error(e);
        this.error = "Network error while saving SFTP";
        this.loading = false;
      }
    },
    onSftpModalConnected(data) {
      this.sftpModalShown = false;
      if (this.sftpPromiseResolve)
        this.sftpPromiseResolve({
          host: this.parseHost(this.hostInput).host,
          port: Number(this.parseHost(this.hostInput).port),
          username: this.sftpUser,
          password: this.sftpPassword,
          path: this.libraryPath,
        });
    },

    onSftpModalCancel() {
      this.sftpModalShown = false;
      if (this.sftpPromiseReject) this.sftpPromiseReject("cancelled");
      this.loginError = "You must set up SFTP before continuing";
      setTimeout(() => (this.loginError = ""), 4000);
    },
    async hashPassword(password) {
      const encoder = new TextEncoder();
      const data = encoder.encode(password);
      const hashBuffer = await crypto.subtle.digest("SHA-256", data);
      return Array.from(new Uint8Array(hashBuffer))
        .map((b) => b.toString(16).padStart(2, "0"))
        .join("");
    },

    async checkUsernameExists(username) {
      try {
        const api = getApiBase();
        const res = await fetch(
          `${api}/api/check-username?username=${encodeURIComponent(username)}`,
        );
        if (!res.ok) return false;
        const data = await res.json();
        return data.exists || false;
      } catch {
        return false;
      }
    },

    async submitRegister() {
      try {
        if (this.registerForm.password !== this.registerForm.confirmPassword)
          throw new Error("Passwords do not match");

        const usernameTaken = await this.checkUsernameExists(
          this.registerForm.username,
        );
        if (usernameTaken) throw new Error("Username is already taken");

        const hashed = await this.hashPassword(this.registerForm.password);
        const payload = { ...this.registerForm, password: hashed };

        const api = getApiBase();
        const res = await fetch(`${api}/api/register`, {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(payload),
        });
        const data = await res.json();
        if (!res.ok) throw new Error(data.error || "Registration failed");

        this.loginError = "";
        this.positiveRespone = "Registration successful!";
        setTimeout(() => (this.positiveRespone = ""), 3000);
      } catch (err) {
        this.positiveRespone = "";
        this.loginError = err.message || "Registration failed";
        setTimeout(() => (this.loginError = ""), 3000);
      }
    },

    closeSftpModal() {
      this.sftpModalShown = false;
      this.error = "";
      this.success = "";
      this.hostError = "";
    },
    async onInitServer() {
      if (!this.localSftpFolderName && this.$refs.folderInput) {
        this.waitingForFolderPick = true;
        this.$refs.folderInput.click();
        return;
      }

      this.initServerLoading = true;
      this.error = "";
      this.initServerHighlight = false;
      try {
        const api = getApiBase();
        const envRes = await fetch(`${api}/sftp/env`, {
          credentials: "include",
        });
        const env = await envRes.json();
        console.log("Env:", env);

        if (!env.rclone_installed) {
          const pm = (env.package_managers || []).find((p) => p.installed);
          if (!pm) {
            this.error =
              "rclone is not installed and no supported package manager was detected. Install from https://rclone.org/install/";
            return;
          }

          const installRes = await fetch(`${api}/sftp/install-rclone`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            credentials: "include",
            body: JSON.stringify({ package_manager: pm.name }),
          });
          const installData = await installRes.json().catch(() => ({}));
          if (!installRes.ok) {
            this.error =
              installData.details ||
              installData.error ||
              `Failed to install rclone using ${pm.name}`;
            return;
          }
        }

        const localPort = this.defaultSftpPort || 2222;
        const startBody = {
          folder_name: this.localSftpFolderName,
          port: localPort,
          user: this.sftpUser || "FlacPlayerUser",
          pass: this.sftpPassword,
        };

        const startRes = await fetch(`${api}/sftp/start-local`, {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          credentials: "include",
          body: JSON.stringify(startBody),
        });
        const startData = await startRes.json().catch(() => ({}));

        if (!startRes.ok) {
          console.error("start-local error:", startData);
          this.error =
            startData.details ||
            startData.error ||
            "Failed to start local SFTP server";
          return;
        }

        console.log("Local SFTP folder actually served:", startData.folder);

        this.localSftpFolderPath = startData.folder;
        this.showInitServer = false;
        this.hostInput = `localhost:${localPort}`;

        console.log("submitSftpCreds");
        await this.submitSftpCreds();
        this.error = "Error while initializing local server";
      } finally {
        this.initServerLoading = false;
        this.waitingForFolderPick = false;
      }
    },
    async loginAsGuest() {
      try {
        this.loginForm.username = "guest";
        this.sftpUser = "guest";
        this.libraryPath = "guest/library";
        this.hostInput = "localhost:2222";

        const guestUsername = "guest";
        const guestHash =
          "84983c60f7daadc1cb8698621f802c0d9f9a3c3c295c810748fb048115c186ec";

        const api = getApiBase();
        const res = await fetch(`${api}/login`, {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          credentials: "include",
          body: JSON.stringify({
            username: guestUsername,
            password: guestHash,
          }),
        });

        if (!res.ok) {
          const data = await res.json();
          this.loginError = data.error || "Guest login failed";
          setTimeout(() => (this.loginError = ""), 3000);
          return;
        }

        console.log("Guest login successful.");
        this.showLoginPopup = false;
        this.positiveRespone = "Logged in as Guest";
        setTimeout(() => (this.positiveRespone = ""), 3000);

        const ok = await this.ensureSftpConnected(guestUsername);
        console.log("ensureSftpConnected(guest) =", ok);
        if (!ok) return;

        console.log("reached route");
        this.$router.push("/library");
      } catch (err) {
        console.error("Guest login error:", err);
        this.loginError = "Failed to connect to server";
        setTimeout(() => (this.loginError = ""), 3000);
      }
    },
    async submitLogin() {
      try {
        const hashed = await this.hashPassword(this.loginForm.password);
        const api = getApiBase();
        const res = await fetch(`${api}/login`, {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          credentials: "include",
          body: JSON.stringify({
            username: this.loginForm.username,
            password: hashed,
            rememberMe: this.loginForm.rememberMe,
          }),
        });

        if (!res.ok) {
          const data = await res.json();
          throw new Error(data.error || "Login failed");
        }

        this.showLoginPopup = false;

        const ok = await this.ensureSftpConnected(this.loginForm.username);
        if (!ok) return;
        this.$router.push("/library");
      } catch (err) {
        this.loginError = err.message || "Login failed";
        setTimeout(() => (this.loginError = ""), 3000);
      }
    },
    closeLoginPopup() {
      this.showLoginPopup = false;
      this.loginError = "";
    },
    async ensureSftpConnected(username) {
      try {
        const api = getApiBase();
        const statusRes = await fetch(`${api}/api/sftp/status`, {
          credentials: "include",
        });
        const status = statusRes.ok ? await statusRes.json() : {};

        let creds;

        if (status.status === "missing") {
          this.sftpUser = username;
          this.libraryPath = `${username}/library`;
          this.hostInput = `localhost:${this.defaultSftpPort || 22}`;
          this.sftpExists = false;
          this.sftpModalShown = true;

          const input = await new Promise((resolve, reject) => {
            this.sftpPromiseResolve = resolve;
            this.sftpPromiseReject = reject;
          });

          creds = {
            host: input.host,
            port: Number(input.port),
            username: input.username,
            password: input.password,
            path: input.path,
          };
        } else {
          this.sftpUser = status.username;
          this.libraryPath = status.path || `${username}/library`;
          this.hostInput = status.host + (status.port ? `:${status.port}` : "");
          this.sftpExists = true;
          this.sftpModalShown = true;

          const input = await new Promise((resolve, reject) => {
            this.sftpPromiseResolve = resolve;
            this.sftpPromiseReject = reject;
          });

          creds = {
            host: status.host,
            port: Number(status.port),
            username: status.username,
            password: input.password,
            path: this.libraryPath,
          };
        }

        const res = await fetch(`${api}/sftp/creds`, {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          credentials: "include",
          body: JSON.stringify(creds),
        });

        const data = await res.json().catch(() => ({}));

        if (!res.ok || !data.connected) {
          this.loginError =
            data.details || data.error || "SFTP connection failed";
          setTimeout(() => (this.loginError = ""), 4000);
          return false;
        }

        return true;
      } catch (err) {
        console.error("ensureSftpConnected error:", err);
        this.loginError = "Failed to connect to SFTP server";
        setTimeout(() => (this.loginError = ""), 4000);
        return false;
      }
    },
  },
};
</script>

<style scoped>
.input-field {
  @apply border border-neutral-300 dark:border-neutral-700  placeholder-neutral-500 dark:placeholder-neutral-400 focus:ring-2 focus:ring-neutral-500 focus:border-transparent transition-all w-full px-4 py-3 rounded-lg;
}

.checkbox-field {
  @apply border-neutral-300 dark:border-neutral-700 rounded focus:ring-neutral-500 focus:ring-2 mt-0.5 w-4 h-4 text-neutral-600 bg-white dark:bg-neutral-800;
}
</style>
<style scoped>
.slide-right-enter-from {
  transform: translateX(100%);
  opacity: 0;
}
.slide-right-enter-to {
  transform: translateX(0);
  opacity: 1;
}
.slide-right-enter-active {
  transition:
    transform 0.3s ease,
    opacity 0.3s ease;
}

.slide-right-leave-from {
  transform: translateX(0);
  opacity: 1;
}
.slide-right-leave-to {
  transform: translateX(100%);
  opacity: 0;
}
.slide-right-leave-active {
  transition:
    transform 0.3s ease,
    opacity 0.3s ease;
}
@keyframes dock-bounce {
  0% {
    transform: translateY(0);
  }
  20% {
    transform: translateY(-6px);
  }
  40% {
    transform: translateY(0);
  }
  60% {
    transform: translateY(-3px);
  }
  80% {
    transform: translateY(0);
  }
  100% {
    transform: translateY(-1px);
  }
}

.dock-bounce {
  animation: dock-bounce 0.6s ease-out;
  animation-iteration-count: 3;
}
</style>
