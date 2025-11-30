<template>
  <div class="fixed">
    <div
      class="bg-stone-100 dark:bg-black font-sans relative md:fixed overflow-auto w-screen overflow-hidden"
    >
      <button
        @click="showSidebar = !showSidebar"
        class="md:hidden fixed top-12 left-1/2 -translate-x-1/2 -ml-2 z-40 w-12 h-6 rounded-lg bg-stone-300 dark:bg-neutral-900 text-white hover:bg-stone-400 dark:hover:bg-neutral-800 flex items-center justify-center shadow-lg transition-colors z-50"
        aria-label="Toggle library menu"
        :class="showSidebar ? ' translate-y-[19.9rem]' : '-translate-y-[3rem]'"
      >
        <svg
          v-if="!showSidebar"
          class="w-6 h-6"
          viewBox="0 0 24 24"
          fill="currentColor"
        >
          <path
            d="M12 15.5c-.26 0-.51-.1-.71-.29l-5-5a1.003 1.003 0 011.42-1.42L12 13.09l4.29-4.3a1.003 1.003 0 111.42 1.42l-5 5c-.2.19-.45.29-.71.29z"
          />
        </svg>
        <svg v-else class="w-6 h-6" viewBox="0 0 24 24" fill="currentColor">
          <path
            d="M12 8.5c.26 0 .51.1.71.29l5 5a1.003 1.003 0 11-1.42 1.42L12 10.91l-4.29 4.3a1.003 1.003 0 01-1.42-1.42l5-5c.2-.19.45-.29.71-.29z"
          />
        </svg>
      </button>
      <div class="h-full flex flex-col md:flex-row">
        <!-- Sidebar -->
        <div
          v-show="showSidebar || windowWidth >= 768"
          :class="[
            'relative xl:w-custom-76 md:w-48 w-dvw m-2 mr-0 bg-white/80 dark:bg-stone-900/80 backdrop-blur-xl border-r border-stone-200 dark:border-stone-700 rounded-2xl scroll-smooth transition-all duration-300 ease-out z-30 transform',
            'md:translate-y-0 md:opacity-100 md:static md:top-auto md:left-auto',
          ]"
          class="absolute md:relative top-0 left-0 transition duration-300"
        >
          <div class="mt-4 ml-4 md:mt-0 md:ml-0 md:p-6">
            <div class="flex justify-between mr-8">
              <p
                class="text-xl font-semibold text-stone-900 flex mb-3 md:mb-8 dark:text-white"
              >
                Music
              </p>
              <button
                @click.prevent="$router.push('/profile')"
                class="md:hidden w-10 h-10 rounded-full bg-stone-200 text-white flex items-center justify-center shadow-lg hover:bg-stone-300 transition-colors"
                aria-label="Profile"
              >
                <svg
                  class="w-10 h-10 fill-neutral-800"
                  viewBox="-2 -2 30 30"
                  preserveAspectRatio="xMidYMid meet"
                >
                  <path
                    d="M12.7148 25.4395C19.7363 25.4395 25.4395 19.7461 25.4395 12.7246C25.4395 5.70312 19.7363 0 12.7148 0C5.69336 0 0 5.70312 0 12.7246C0 19.7461 5.69336 25.4395 12.7148 25.4395ZM12.7148 23.623C6.68945 23.623 1.81641 18.75 1.81641 12.7246C1.81641 6.69922 6.68945 1.82617 12.7148 1.82617C18.7402 1.82617 23.6133 6.69922 23.6133 12.7246C23.6133 18.75 18.7402 23.623 12.7148 23.623ZM21.2988 20.9668L21.2695 20.8496C20.5957 19.0332 17.2949 16.9629 12.7148 16.9629C8.1543 16.9629 4.85352 19.0234 4.16016 20.8301L4.13086 20.9668C6.38672 23.2227 10.0684 24.5703 12.7246 24.5703C15.3809 24.5703 19.0039 23.252 21.2988 20.9668ZM12.7148 14.8145C15.1367 14.834 17.0215 12.7637 17.0215 10.0684C17.0215 7.5293 15.1172 5.41992 12.7148 5.41992C10.3125 5.41992 8.39844 7.5293 8.4082 10.0684C8.41797 12.7637 10.293 14.7949 12.7148 14.8145Z"
                    fill-opacity="0.85"
                  />
                </svg>
              </button>
            </div>

            <nav class="space-y-2 hidden md:inline text-lg font-medium">
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
                href="#"
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
                href="#"
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
              <a
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
              </a>
            </nav>

            <!-- Library section -->
            <div class="mt-0 md:mt-4">
              <Check />
              <p
                class="text-lg font-medium text-stone-500 mb-3 dark:text-stone-400"
              >
                Library
              </p>
              <nav class="space-y-2">
                <a
                  v-for="item in libraryNav"
                  :key="item.name"
                  href="#"
                  @click.prevent="
                    setActive(item.name);
                    showSidebar = true;
                  "
                  :class="
                    item.active
                      ? 'text-red-600 bg-red-50 dark:bg-red-900/20 dark:text-red-400'
                      : 'text-stone-700 dark:text-stone-300 hover:bg-stone-100 dark:hover:bg-stone-700'
                  "
                  class="flex items-center px-3 py-2 rounded-lg space-x-3 text-xl"
                >
                  <div class="flex items-center justify-center flex-shrink-0">
                    <svg
                      v-if="
                        item.icon && /<\s*(path|g)\b/i.test(item.icon.trim())
                      "
                      class="block fill-black dark:fill-white w-8 h-8"
                      viewBox="0 0 38 28"
                      preserveAspectRatio="xMidYMid meet"
                      v-html="item.icon"
                    ></svg>
                    <span v-else class="block">{{ item.icon }}</span>
                  </div>

                  <span>{{ item.name }}</span>
                </a>
              </nav>
            </div>
          </div>
          <div
            v-if="activeView === 'Artists'"
            class="md:py-[2rem] h-20 md:h-[30rem] border-2 border-dashed border-stone-500 hover:border-neutral-700 transition-colors bg-neutral-100 text-stone-800 dark:bg-neutral-900 dark:text-stone-400 rounded-3xl m-2 mx-1 xl:mx-6 md:ml-1 md:mr-2 overflow-x-auto overflow-y-scroll"
            style="scrollbar-width: none"
          >
            <!-- Alphabet quick jump for artists -->
            <div
              class="flex flex-row md:flex-col justify-start md:justify-center mr-6 md:mr-2 z-40 md:text-right text-lg px-2 md:pr-0"
            >
              <div
                v-for="(artists, letter) in groupedArtists"
                :key="letter"
                class="md:mb-6 md:text-right"
              >
                <h2 :id="letter" class="text-xl font-bold mb-1 md:text-right">
                  {{ letter }}
                </h2>
                <ul
                  class="m-1 md:pl-4 flex flex-row gap-3 items-center whitespace-nowrap overflow-x-auto md:flex-col md:gap-1 md:items-end md:whitespace-normal md:overflow-visible"
                  style="scrollbar-width: thin; scroll-snap-type: x mandatory"
                >
                  <li
                    v-for="artist in artists"
                    :key="artist.name"
                    class="hover:text-red-600 cursor-pointer md:text-right flex-shrink-0 inline-flex md:inline"
                  >
                    <a @click="selectArtist(artist.name)">{{ artist.name }}</a>
                  </li>
                </ul>
              </div>
            </div>
          </div>
          <div
            v-if="activeView === 'Genres'"
            class="md:py-[2rem] h-20 md:h-[30rem] border-2 border-dashed border-stone-500 hover:border-neutral-700 transition-colors bg-neutral-100 text-stone-800 dark:bg-neutral-900 dark:text-stone-400 rounded-3xl m-2 mx-1 xl:mx-6 md:ml-1 md:mr-2 overflow-x-auto overflow-y-scroll"
            style="scrollbar-width: none"
          >
            <!-- Alphabet quick jump -->
            <div
              class="flex flex-row md:flex-col justify-start md:justify-center mr-6 md:mr-2 z-40 md:text-right text-lg px-2 md:pr-0"
            >
              <div
                v-for="(genres, letter) in groupedGenres"
                :key="letter"
                class="md:mb-6 md:text-right"
              >
                <h2 :id="letter" class="text-xl font-bold mb-1 md:text-right">
                  {{ letter }}
                </h2>
                <ul
                  class="m-1 md:pl-4 flex flex-row gap-3 items-center whitespace-nowrap overflow-x-auto md:flex-col md:gap-1 md:items-end md:whitespace-normal md:overflow-visible"
                  style="scrollbar-width: thin; scroll-snap-type: x mandatory"
                >
                  <li
                    v-for="genre in genres"
                    :key="genre"
                    class="hover:text-red-600 cursor-pointer md:text-right flex-shrink-0 inline-flex md:inline"
                  >
                    <a :href="'#' + genre" class="">{{ genre }}</a>
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>

        <!-- Main content -->
        <div class="flex-1 flex flex-col relative h-screen">
          <div
            class="relative md:absolute flex items-center justify-between h-[4rem] max-w-screen top-0 left-0 right-0 z-[60] px-6 py-4 bg-[linear-gradient(to_bottom,rgba(var(--bg),1)_0%,rgba(var(--bg),1)_40%,rgba(var(--bg),0.9)_55%,rgba(var(--bg),0.93)_70%,rgba(var(--bg),1)_100%)] backdrop-blur-sm dark:[--bg:0,0,0] [--bg:241,241,241]"
          >
            <p
              v-if="activeView === 'Recently Added'"
              class="text-xl font-semibold text-stone-600 dark:text-white"
            >
              Recently Added
            </p>
            <p
              v-if="activeView === 'Albums'"
              class="text-xl font-semibold text-stone-600 dark:text-white"
            >
              Albums
            </p>
            <p
              v-if="activeView === 'Artists'"
              class="text-xl font-semibold text-stone-600 dark:text-white"
            >
              Artists
            </p>
            <p
              v-if="activeView === 'Songs'"
              class="text-xl font-semibold text-stone-600 dark:text-white"
            >
              Songs
            </p>
            <p
              v-if="activeView === 'Genres'"
              class="text-xl font-semibold text-stone-600 dark:text-white"
            >
              Genres
            </p>
            <div
              class="items-center w-30 flex space-x-4 bg-stone-200 focus-within:ring-2 focus-within:ring-red-500 dark:bg-stone-700 text-stone-900 rounded-[12rem] p-1 px-2"
            >
              <svg
                class="w-8 h-6 fill-stone-500 dark:fill-stone-400"
                viewBox="0 0 18 28"
                preserveAspectRatio="xMidYMid meet"
              >
                <path
                  d="M0 10.2337C0 15.875 4.58576 20.4639 10.2337 20.4639C12.3037 20.4639 14.2165 19.8282 15.8074 18.7534L21.5753 24.5249C21.9374 24.8968 22.4376 25.0765 22.9506 25.0765C24.0575 25.0765 24.8515 24.2297 24.8515 23.1526C24.8515 22.6391 24.6717 22.1683 24.31 21.7999L18.5839 16.0506C19.766 14.4211 20.4706 12.4075 20.4706 10.2337C20.4706 4.58576 15.8817 0 10.2337 0C4.58576 0 0 4.58576 0 10.2337ZM2.60873 10.2337C2.60873 6.02948 6.02948 2.60873 10.2337 2.60873C14.4411 2.60873 17.8521 6.02948 17.8521 10.2337C17.8521 14.4313 14.4411 17.8521 10.2337 17.8521C6.02948 17.8521 2.60873 14.4313 2.60873 10.2337Z"
                  fill-opacity="0.85"
                />
              </svg>
              <label
                class="relative h-full flex-1 whitespace-nowrap overflow-hidden w-20 md:w-[12rem] lg:w-[14rem] dark:text-white placeholder-stone-500 dark:placeholder-stone-400 focus:outline-none bg-stone-200 dark:bg-stone-700 text-transparent caret-white rounded-2xl"
              >
                <input
                  type="text"
                  v-model="searchQuery"
                  class="absolute inset-0 w-full h-4 bg-transparent border-none outline-none text-transparent caret-white"
                />
                <template v-if="!displayTokens.length">
                  <span class="text-stone-500 dark:text-stone-300"
                    >Find in Albums</span
                  >
                </template>
                <template v-else>
                  <span
                    v-for="(t, idx) in displayTokens"
                    :key="idx"
                    class="flex items-center text-base"
                  >
                    <template v-if="t.isTag && t.value">
                      <span
                        class="rounded-2xl pl-2 pr-1 text-sm font-semibold -translate-x-1"
                        :class="t.class"
                      >
                        {{ t.label }}
                      </span>
                      <span class="text-white -translate-x-1">{{
                        t.value
                      }}</span>
                    </template>
                    <template v-else>
                      <span :class="t.class">{{ t.value || t.raw }}</span>
                    </template>
                  </span>
                </template>
              </label>
            </div>
          </div>

          <!-- Dropzone -->
          <div
            class="library-dropzone w-full h-full z-50"
            @dragover.prevent="onDragOver"
            @dragenter.prevent="onDragEnter"
            @dragleave.prevent="onDragLeave"
            @drop.prevent="handleDrop"
          >
            <form ref="uploadForm" class="absolute inset-0 w-full h-full">
              <input
                ref="fileInput"
                type="file"
                multiple
                webkitdirectory
                directory
                accept="audio/flac,.flac,.FLAC"
                style="display: none"
                @change="handleFiles"
              />
              <input
                ref="filepickerInput"
                type="file"
                multiple
                accept="audio/flac,.flac,.FLAC,audio/x-flac"
                style="display: none"
                @change="handleFiles"
              />
            </form>

            <div
              v-if="dragOver || UploadGuide"
              class="absolute inset-0 bg-black/30 dark:bg-white/20 flex items-center justify-center z-40 rounded-lg pointer-events-none"
              @click="hideUploadGuide"
            >
              <p
                class="text-white dark:text-black font-bold text-lg bg-neutral-900 dark:bg-neutral-200 p-2 rounded-lg"
              >
                Drag and drop files to upload
              </p>
              <p
                class="fixed bottom-6 right-20 z-50 text-white dark:text-black font-bold text-lg bg-neutral-900 dark:bg-neutral-200 p-2 rounded-lg"
              >
                Or add Files with button &#10551;
              </p>
            </div>
            <!-- ARTIST VIEW -->
            <div
              v-else-if="activeView === 'Artists'"
              class="relative md:overflow-y-auto h-dvh pt-[3rem] content-start"
            >
              <div class="p-8 space-y-8">
                <h1
                  class="text-3xl font-bold mb-4 text-stone-800 dark:text-stone-100"
                >
                  {{ selectedArtist ? selectedArtist : "" }}
                </h1>
                <div v-if="selectedArtist">
                  <button
                    v-if="selectedArtist"
                    @click="selectedArtist = null"
                    class="text-sm text-blue-500 hover:underline mb-2"
                  >
                    ← Back to Artists
                  </button>
                  <div
                    class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-5 gap-6"
                  >
                    <div
                      v-for="album in artistAlbums"
                      :key="album.id"
                      class="group cursor-pointer"
                      @click="playFromSong(song)"
                    >
                      <RouterLink :to="`/album/${album.id}`">
                        <div
                          class="aspect-square mb-3 rounded-xl relative overflow-hidden"
                        >
                          <img
                            :alt="album.name"
                            :src="album.cover"
                            class="object-cover group-hover:scale-105 transition-transform duration-300 w-full h-full"
                          />
                          <div
                            v-if="
                              album.songs?.length &&
                              album.songs.every((s) => s.missing)
                            "
                            class="absolute inset-0 bg-neutral-900/55 backdrop-blur-[1px] text-white flex items-center justify-center text-center px-3 pointer-events-none"
                          >
                            <p class="text-xs text-stone-200 shadow-lg">
                              Unavailable
                            </p>
                          </div>
                        </div>
                        <p class="font-medium text-stone-900 dark:text-white">
                          {{ album.name }}
                        </p>
                        <p class="text-sm text-stone-500 dark:text-stone-400">
                          {{ album.artist }}
                        </p>
                      </RouterLink>
                    </div>
                  </div>
                </div>
                <div
                  v-else
                  class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-5 gap-6"
                >
                  <div
                    v-for="artist in filteredArtists"
                    :key="artist.name"
                    class="group relative cursor-pointer rounded-xl overflow-hidden aspect-square"
                    @click="selectArtist(artist.name)"
                  >
                    <div class="absolute inset-0">
                      <img
                        :src="artist.randomCover"
                        alt=""
                        class="w-full h-full object-cover scale-110 blur-sm brightness-75"
                      />
                    </div>
                    <div
                      class="absolute inset-0 bg-black/40 backdrop-blur-sm flex flex-col items-center justify-center text-center transition-all duration-300 group-hover:bg-black/30"
                    >
                      <h2
                        class="text-2xl font-semibold text-white drop-shadow-md"
                      >
                        {{ artist.name }}
                      </h2>
                      <p class="text-lg text-stone-300 mt-1">
                        {{ artist.albumCount }} albums
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <!-- ALBUMView -->
            <div
              v-if="activeView === 'Albums' || activeView === 'Recently Added'"
            >
              <div
                class="grid grid-cols-2 Umd-60:grid-cols-3 Ulg-70:grid-cols-4 Uxl:grid-cols-5 U2xl:grid-cols-6 gap-6 p-8 pb-0 relative overflow-y-auto h-dvh md:pt-[6rem] content-start"
              >
                <template v-if="displayAlbums.length">
                  <div
                    v-for="album in displayAlbums"
                    :key="album.id"
                    class="group cursor-pointer max-h-1/4"
                  >
                    <RouterLink :to="`/album/${album.id}`">
                      <div
                        class="aspect-square mb-3 rounded-xl relative overflow-hidden"
                      >
                        <img
                          :alt="album.name"
                          :src="album.cover"
                          class="object-cover group-hover:scale-105 transition-transform duration-300 w-full h-full"
                        />
                        <div
                          v-if="
                            album.songs?.length &&
                            album.songs.every((s) => s.missing)
                          "
                          class="absolute inset-0 bg-neutral-900/55 backdrop-blur-[1px] text-white flex items-center justify-center text-center px-3 pointer-events-none"
                        >
                          <p
                            class="text-xs text-stone-200 shadow-lg bg-black/50 p-2 rounded-lg"
                          >
                            Unavailable
                          </p>
                        </div>
                      </div>
                      <p class="font-medium text-stone-900 dark:text-white">
                        {{ album.name }}
                      </p>
                      <p class="text-sm text-stone-500 dark:text-stone-400">
                        {{ album.artist }}
                      </p>
                    </RouterLink>
                  </div>
                </template>
                <template v-else>
                  <div class="w-full">
                    <p
                      class="text-stone-500 dark:text-stone-400 col-span-full text-left"
                    >
                      No albums found
                    </p>
                  </div>
                </template>
              </div>
            </div>

            <!-- GENREView -->
            <div
              v-else-if="activeView === 'Genres'"
              class="relative md:overflow-y-auto h-dvh pt-[3rem] content-start"
            >
              <div class="p-8 space-y-8">
                <!-- If searching, show flat filtered albums -->
                <template v-if="searchQuery && filteredGenres.length">
                  <div
                    v-for="group in filteredGenres"
                    :key="group.genre"
                    class="space-y-4"
                  >
                    <div
                      class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-5 gap-6"
                    >
                      <div
                        v-for="album in group.albums"
                        :key="album.id"
                        class="group cursor-pointer"
                      >
                        <RouterLink :to="`/album/${album.id}`">
                          <div
                            class="aspect-square mb-3 rounded-xl relative overflow-hidden"
                          >
                            <img
                              :alt="album.name"
                              :src="album.cover"
                              class="object-cover group-hover:scale-105 transition-transform duration-300 w-full h-full"
                            />
                            <div
                              v-if="
                                album.songs?.length &&
                                album.songs.every((s) => s.missing)
                              "
                              class="absolute inset-0 bg-neutral-900/55 backdrop-blur-[1px] text-white flex items-center justify-center text-center px-3 pointer-events-none"
                            >
                              <div class="space-y-1">
                                <p
                                  class="text-xs text-stone-200 shadow-lg bg-black/50 p-2 rounded-lg"
                                >
                                  Unavailable
                                </p>
                              </div>
                            </div>
                          </div>
                          <p class="font-medium text-stone-900 dark:text-white">
                            {{ album.name }}
                          </p>
                          <p class="text-sm text-stone-500 dark:text-stone-400">
                            {{ album.artist }}
                          </p>
                        </RouterLink>
                      </div>
                    </div>
                  </div>
                </template>
                <!-- If no search, show grouped by genre -->
                <template v-else>
                  <div
                    v-for="(group, idx) in groupedByGenre"
                    :key="group.genre"
                    class="space-y-4"
                  >
                    <span
                      :id="group.genre"
                      class="scroll-mt-[calc(50vh-2rem)]"
                    ></span>
                    <h2
                      class="text-2xl font-semibold text-stone-800 dark:text-stone-100"
                    >
                      {{ group.genre }}
                    </h2>

                    <div
                      class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-5 gap-6"
                    >
                      <div
                        v-for="album in group.albums"
                        :key="album.id"
                        class="group cursor-pointer"
                      >
                        <RouterLink :to="`/album/${album.id}`">
                          <div
                            class="aspect-square mb-3 rounded-xl overflow-hidden relative"
                          >
                            <img
                              :src="album.cover"
                              :alt="album.name"
                              class="object-cover w-full h-full transition-transform group-hover:scale-105"
                            />
                            <div
                              v-if="
                                album.songs?.length &&
                                album.songs.every((s) => s.missing)
                              "
                              class="absolute inset-0 bg-neutral-900/55 backdrop-blur-[1px] text-white flex items-center justify-center text-center px-3 pointer-events-none"
                            >
                              <p
                                class="text-xs text-stone-200 shadow-lg bg-black/50 p-2 rounded-lg"
                              >
                                Unavailable
                              </p>
                            </div>
                          </div>
                          <p class="font-medium text-stone-900 dark:text-white">
                            {{ album.name }}
                          </p>
                          <p class="text-sm text-stone-500 dark:text-stone-400">
                            {{ album.artist }}
                          </p>
                        </RouterLink>
                      </div>
                    </div>
                    <div
                      v-if="idx !== groupedByGenre.length - 1"
                      class="border-t border-stone-300 dark:border-stone-700 mt-4"
                    ></div>
                  </div>
                </template>
              </div>
            </div>
            <!-- SongVIEW -->
            <div
              v-if="activeView === 'Songs'"
              class="p-6 relative md:overflow-y-auto h-dvh pt-[5rem] w-full content-start"
            >
              <div
                class="overflow-x-auto rounded-2xl shadow-sm border border-stone-200 dark:border-stone-700 p-2"
              >
                <table class="table-auto mx-auto text-md w-full">
                  <thead
                    class="bg-stone-100 dark:bg-stone-950 text-stone-700 dark:text-stone-300 uppercase tracking-wide text-lg"
                  >
                    <tr>
                      <th
                        v-for="(col, index) in visibleColumns"
                        :key="col.key"
                        :class="[
                          'px-4 py-3 whitespace-nowrap cursor-pointer select-none text-left truncate',
                          col.key === 'duration' ? 'w-[6em]' : 'relative',
                          col.key === 'album' || col.key === 'genre'
                            ? 'hidden lg:table-cell'
                            : '',
                        ]"
                        draggable="true"
                        @dragstart="onDragStart(index)"
                        @dragover.prevent
                        @drop="onDrop(index)"
                        @click="sortBy(col.key)"
                      >
                        <span class="inline-flex items-center gap-1">
                          {{ col.label }}
                          <span v-if="sortKey === col.key">
                            <span
                              v-if="sortOrder === 'asc'"
                              class="inline-block"
                              >⇧</span
                            >
                            <span
                              v-else
                              class="inline-block transform rotate-180 translate-y-[0.125rem]"
                              >⇧</span
                            >
                          </span>
                        </span>
                      </th>
                      <th></th>
                    </tr>
                  </thead>

                  <tbody
                    class="divide-y divide-stone-100 dark:divide-stone-700 text-stone-900 dark:text-stone-200 text-[1.1rem]"
                  >
                    <tr
                      v-for="(song, i) in sortedSongs"
                      :key="song.id"
                      :class="[
                        'dark:even:bg-stone-800 dark:odd:bg-stone-900 even:bg-neutral-50 odd:bg-white select-none hover:bg-stone-100 dark:hover:bg-stone-700/40 transition-colors',
                        song.missing
                          ? 'opacity-60 text-stone-400 dark:text-stone-500 pointer-events-none'
                          : '',
                      ]"
                      :title="[
                        song.missing
                          ? 'Not found on your current SFTP server'
                          : '',
                      ]"
                      @click="playFromSong(song)"
                    >
                      <td
                        v-for="col in visibleColumns"
                        :key="col.key"
                        :class="[
                          'px-4 py-2 truncate border-2 border-neutral-200 dark:border-neutral-700 4xl:border-0',
                          col.key === 'duration'
                            ? 'w-[6rem]'
                            : 'relative max-w-0',
                          col.key === 'album' || col.key === 'genre'
                            ? 'hidden lg:table-cell'
                            : '',
                          'cursor-pointer',
                        ]"
                      >
                        <div
                          v-if="col.key !== 'duration'"
                          class="absolute px-4 py-2 inset-0 w-full overflow-hidden whitespace-nowrap text-ellipsis"
                        >
                          {{ song[col.key] }}
                        </div>
                        <span v-else>{{ song[col.key] }}</span>
                      </td>
                      <td class="w-0 4xl:w-72"></td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Upload popup -->
    <div
      v-if="uploadingSongs.length"
      class="fixed inset-0 bg-black/50 flex items-center justify-center z-50"
      @click.self="clearUploading"
    >
      <div
        class="bg-white dark:bg-stone-800 rounded-lg shadow-lg p-6 w-[90vw] max-w-screen-xl max-h-[80vh] overflow-y-auto z-[60]"
      >
        <!-- Actions -->
        <div class="flex justify-end">
          <button
            @click="clearUploading"
            class="px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600"
          >
            Close
          </button>
        </div>
        <h2 class="text-lg font-semibold text-stone-900 dark:text-white w-1/2">
          Uploading & Parsing
        </h2>

        <!-- Two-column layout -->
        <div class="grid grid-cols-2 gap-6 max-h-[60vh]">
          <!-- Column 1: Upload progress -->
          <div>
            <h3 class="text-stone-700 dark:text-stone-300 font-medium mb-2">
              Upload Progress
            </h3>
            <ul class="space-y-2">
              <li
                v-for="(song, idx) in uploadingSongs"
                :key="idx"
                class="flex items-center space-x-2"
              >
                <span v-if="!song.done" class="animate-spin">⏳</span>
                <span v-else class="text-green-500">✅</span>
                <span class="text-stone-900 dark:text-stone-200">
                  {{ song.artist || "Unknown Artist" }} –
                  {{ song.title || "Parsing..." }}
                </span>
              </li>
            </ul>
          </div>
          <div v-if="uploadTree" class="mt-4 h-full">
            <h3 class="text-stone-700 dark:text-stone-300 font-medium mb-2">
              Library Tree
            </h3>
            <!-- prettier-ignore -->
            <pre class="p-3 bg-stone-100 dark:bg-stone-700 text-stone-900 dark:text-stone-200 rounded-lg text-sm font-mono overflow-auto h-full leading-tight whitespace-pre tracking-normal">{{ uploadTree }}</pre>
          </div>
        </div>
      </div>
    </div>
    <div
      class="fixed bottom-4 right-4 z-50 w-14 h-14 flex items-center justify-center bg-stone-800 rounded-full shadow-lg hover:bg-stone-700 transition-colors cursor-help"
      @click="triggerFilePicker(userId)"
    >
      <div class="relative w-7 h-7">
        <div
          class="absolute inset-0 w-full h-1 bg-white rounded-full transform translate-y-3"
        ></div>
        <div
          class="absolute inset-0 h-full w-1 bg-white rounded-full transform translate-x-3"
        ></div>
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref, computed, onMounted, watch } from "vue";
import { useRouter, useRoute } from "vue-router";
import Fuse from "fuse.js";
import { getApiBase } from "../apiBase";
const router = useRouter();
const route = useRoute();
const username = ref("");
const albums = ref([]);
const sortKey = ref("title");
const sortOrder = ref("asc");
const userId = ref(null);
const showSidebar = ref(true);
const windowWidth = ref(
  typeof window !== "undefined" ? window.innerWidth : 1024,
);
const songs = ref([]);
const dragOver = ref(false);
const searchQuery = ref("");
const UploadGuide = ref(false);
const uploadingSongs = ref([]);
const selectedArtist = ref(null);
const uploadTree = ref("");
const fileInput = ref(null);
const filepickerInput = ref(null);
let progressInterval = null;

import { usePlayer } from "../player/usePlayer";
const player = usePlayer();

function toPlayerTrack(song) {
  return {
    id: song.song_id || song.id,
    title: song.title,
    artist: song.artist || song.album_artist || "Unknown Artist",
    album: song.album || song.album_name || "Unknown Album",
    durationLabel: song.duration,
    cover: song.cover || song.album_cover || null,
    streamUrl: `${getApiBase()}/stream/${song.song_id || song.id}`,
  };
}

function shuffleList(list) {
  const arr = list.slice();
  for (let i = arr.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1));
    [arr[i], arr[j]] = [arr[j], arr[i]];
  }
  return arr;
}

function playFromSong(song) {
  const list = sortedSongs.value || songs.value || [];
  const targetId = song.song_id || song.id;
  const idx = list.findIndex((s) => (s.song_id || s.id) === targetId);
  if (idx === -1) return;

  const below = list.slice(idx).map(toPlayerTrack);
  const above = list.slice(0, idx).map(toPlayerTrack);

  let queue;
  if (player.state.isShuffle) {
    queue = [...shuffleList(below), ...shuffleList(above)];
  } else {
    queue = [...below, ...above];
  }

  if (!queue.length) return;
  player.playNow(queue[0], queue);
}

const libraryNav = ref([
  {
    name: "Recently Added",
    icon: '<path d="M13.0005 26.0077C20.1804 26.0077 26.0077 20.184 26.0077 13.0072C26.0077 5.82734 20.1804 0 13.0005 0C5.82068 0 0 5.82734 0 13.0072C0 20.184 5.82068 26.0077 13.0005 26.0077ZM13.0005 23.4678C7.21737 23.4678 2.53996 18.7904 2.53996 13.0072C2.53996 7.22403 7.21737 2.54352 13.0005 2.54352C18.7837 2.54352 23.4611 7.22403 23.4611 13.0072C23.4611 18.7904 18.7837 23.4678 13.0005 23.4678Z" fill-opacity="0.85"/><path d="M6.61848 14.6032L12.9908 14.6032C13.5447 14.6032 13.973 14.175 13.973 13.6242L13.973 5.33244C13.973 4.78831 13.5447 4.36001 12.9908 4.36001C12.4533 4.36001 12.0183 4.78831 12.0183 5.33244L12.0183 12.6486L6.61848 12.6486C6.06458 12.6486 5.64294 13.08 5.64294 13.6242C5.64294 14.175 6.06458 14.6032 6.61848 14.6032Z" fill-opacity="0.85"/>',
    active: false,
  },
  {
    name: "Artists",
    icon: '<path d="M13.0655 26.0293C13.7761 26.0293 14.3251 25.4803 14.3313 24.7795L14.353 15.5872L18.0711 12.1337C19.5573 12.3055 21.1069 11.6779 22.4413 10.328L15.2189 3.09896C13.8624 4.43644 13.2747 5.94611 13.4562 7.442L1.75568 20.0166C1.24656 20.5715 1.22968 21.2444 1.85378 21.9461L0.137279 24.1083C-0.0646527 24.3582-0.0460203 24.7172 0.207845 24.9777L0.578939 25.3422C0.83591 25.5991 1.18163 25.6244 1.46079 25.4065L3.61053 23.6767C4.28563 24.3106 4.98209 24.2839 5.52362 23.7779L11.7929 17.9631L11.7929 24.7795C11.7929 25.4803 12.3549 26.0293 13.0655 26.0293ZM3.25178 21.0597L14.213 9.50983C14.4221 9.86719 14.6835 10.2112 14.9881 10.5491C15.2957 10.8767 15.6397 11.155 15.974 11.3637L4.4903 22.3049ZM16.5416 1.77627L23.764 9.00839C26.0354 6.7685 26.1916 3.82872 23.9447 1.59203C21.7146-0.625139 18.7815-0.492015 16.5416 1.77627Z" fill-opacity="0.85"/></g>',
    active: false,
  },
  {
    name: "Albums",
    icon: '  <g transform="scale(0.45) translate(0, -10)"><path viewBox="0 0 100 100" d="M 29.06428 52.54882 C 34.703152 52.541927 39.272644 47.972435 39.27953 42.333565 C 39.27953 35.934532 33.369568 32.23077 29.280256 32.118324 L 29.06428 32.118324 C 23.425409 32.125214 18.855917 36.694706 18.849031 42.333565 C 18.840107 47.968636 23.429203 52.54882 29.06428 52.54882 M 30.270903 34.228127 C 30.290537 34.181725 30.299461 34.124599 30.31731 34.076405 C 30.354799 34.085339 30.383354 34.085339 30.420837 34.096043 C 30.36372 34.142445 30.31731 34.171017 30.269117 34.228127 M 28.441334 33.963959 C 27.651262 36.584103 27.241894 39.304211 27.225786 42.040836 C 27.215927 44.977737 27.670759 47.897713 28.57342 50.692478 C 24.138844 50.434624 20.672785 46.766697 20.666109 42.324646 C 20.675339 37.934338 24.063232 34.291332 28.441334 33.963959"/><path d="M 83.103531 42.04977 C 83.103531 26.065567 69.164879 15.491547 56.648842 14.615143 L 56.648842 13.342476 L 0.215979 13.342476 L 0.215979 70.699936 L 56.648842 70.699936 L 56.648842 69.512955 C 71.351448 68.966751 83.103531 56.884449 83.103531 42.04977 M 56.648842 41.36256 C 56.98798 41.551758 57.241436 41.842709 57.309269 42.190773 C 57.393166 42.642365 57.148624 43.190346 56.648842 43.500919 Z M 56.648842 44.284515 C 57.591293 43.916813 58.176754 42.992218 57.996471 42.060474 C 57.875095 41.400051 57.336044 40.853844 56.648842 40.60931 L 56.648842 35.877422 C 59.456558 36.546764 61.634197 38.590542 62.105423 41.060902 C 62.671246 44.029278 60.616779 47.733036 56.648842 48.882545 Z M 54.829979 68.872154 L 2.052688 68.872154 L 2.052688 15.161331 L 54.829979 15.161331 Z M 56.648842 51.643845 L 56.648842 50.606796 C 61.352173 49.47514 64.593636 45.035976 63.763638 40.73069 C 63.140686 37.498146 60.28656 34.849293 56.657764 34.124599 L 56.648842 34.181725 L 56.648842 33.135746 C 61.358669 33.661911 64.921341 37.642643 64.923851 42.381756 C 64.934525 47.129265 61.367455 51.122593 56.648842 51.64563"/>  </g>',
    active: true,
  },
  {
    name: "Songs",
    icon: ' <g> <path d="M15.6495 6.38218L15.6495 1.71466C15.6495 0.953027 15.0241 0.440392 14.2784 0.5984L7.87669 1.99399C7.01029 2.17818 6.54469 2.66598 6.54469 3.42896L6.56467 17.2908C6.61084 17.852 6.35068 18.2076 5.86767 18.3061L3.84709 18.7292C1.25838 19.2618 0 20.5952 0 22.6037C0 24.634 1.53803 26.0358 3.74097 26.0358C5.87768 26.0358 8.5202 24.4437 8.5202 20.9054L8.5202 9.42771C8.5202 8.87054 8.64184 8.72937 9.14663 8.61394L14.8501 7.36529C15.337 7.26187 15.6495 6.88195 15.6495 6.38218Z" fill-opacity="0.85"/></g>',
    active: false,
  },
  {
    name: "Genres",
    icon: '<path d="M29.7786 1.39853L29.9708 4.54482C29.9837 4.80359 29.8767 4.95538 29.6047 5.14314C29.3205 5.33579 29.2744 5.4703 29.2744 5.95993L29.2744 13.5779C31.3802 13.8465 32.6559 14.9824 32.6559 16.7223C32.6559 17.8098 32.2719 18.5685 32.2719 19.4065C32.2719 20.5606 33.9835 21.1537 33.9835 23.7797C33.9835 26.4355 31.7459 28.1204 28.2211 28.1204C26.201 28.1204 24.6035 27.5661 23.6115 26.5899C23.9371 26.1569 24.1967 25.6713 24.3777 25.1364C24.9985 26.071 26.3391 26.6029 28.2211 26.6029C30.884 26.6029 32.4629 25.538 32.4629 23.7797C32.4629 21.7414 30.7451 20.9508 30.7451 19.3741C30.7451 18.4828 31.129 17.5555 31.129 16.7223C31.129 15.4501 29.9931 14.9854 28.2211 14.9854C26.4491 14.9854 25.3131 15.4501 25.3131 16.7223C25.3131 17.5555 25.7002 18.4828 25.7002 19.3741C25.7002 20.2997 25.1071 20.9544 24.6182 21.7508C24.4904 21.0301 24.263 20.457 24.0237 19.9945C24.1171 19.8067 24.1734 19.6159 24.1734 19.4065C24.1734 18.6027 23.8172 17.8719 23.7891 16.853C23.9174 16.4766 24.0166 16.0397 24.0266 15.5172C24.4869 14.4608 25.593 13.7782 27.1611 13.5779L27.1611 5.95993C27.1611 5.47696 27.1216 5.34556 26.8308 5.14314C26.5685 4.95538 26.4615 4.79738 26.4713 4.54482L26.66 1.39853C26.6826 1.1016 26.9023 0.888101 27.1699 0.888101L29.2589 0.888101C29.5398 0.888101 29.756 1.1016 29.7786 1.39853ZM30.3451 23.4641C30.3451 23.7717 30.0925 24.0207 29.7818 24.0207L26.6439 24.0207C26.3301 24.0207 26.0806 23.7717 26.0806 23.4641C26.0806 23.1597 26.3399 22.8942 26.6439 22.8942L29.7818 22.8942C30.0925 22.8942 30.3451 23.1499 30.3451 23.4641ZM29.5865 19.0421C29.5865 19.7932 28.9752 20.4173 28.2211 20.4173C27.4603 20.4173 26.8393 19.7963 26.8393 19.0421C26.8393 18.2782 27.47 17.6572 28.2211 17.6572C28.9752 17.6572 29.5865 18.2813 29.5865 19.0421Z" fill-opacity="0.85"/><path d="M18.6502 1.10515L19.7435 2.05197C19.941 2.23351 20.0475 2.47941 20.0315 2.73507L19.7776 5.36909C19.7385 5.74905 19.6413 5.91107 19.3524 6.0269C19.1531 6.10147 19.1402 6.1503 19.1402 6.46595L19.1304 14.2513C19.1304 14.7498 19.2148 15.2058 19.4931 15.2058C19.684 15.2058 19.8921 14.9745 20.1163 14.6602C20.419 14.2314 20.8723 13.7923 21.6278 13.7923C22.5182 13.7923 23.1617 14.4661 23.1617 15.4129C23.1617 16.7024 22.2904 17.4654 22.2904 18.525C22.2904 19.4305 23.8622 20.3885 23.8622 23.0314C23.8622 26.1363 21.5594 28.1888 18.0864 28.1888C17.0588 28.1888 16.1286 28.0072 15.3326 27.665C15.5477 27.1771 15.5979 26.6029 15.4677 26.0267C16.1509 26.4343 17.0374 26.6517 18.0864 26.6517C20.7169 26.6517 22.3252 25.2846 22.3252 23.0314C22.3252 20.7067 20.7791 19.9361 20.7791 18.5281C20.7791 17.2049 21.6327 16.4219 21.6327 15.4777C21.6327 15.3724 21.5572 15.2903 21.4587 15.2903C20.9695 15.2903 20.6463 16.69 19.4447 16.69C18.8437 16.69 18.3519 16.3295 18.1526 15.5913L18.0301 15.5913C17.8405 16.326 17.3523 16.6962 16.7415 16.6962C15.5265 16.6962 15.2034 15.2903 14.7142 15.2903C14.6223 15.2903 14.5468 15.3724 14.5468 15.4777C14.5468 16.4219 15.3969 17.2049 15.3969 18.5281C15.3969 19.7849 14.1626 20.5338 13.9004 22.3365L13.0016 20.3088C13.4237 19.5233 13.8856 19.0164 13.8856 18.525C13.8856 17.4654 13.0111 16.7024 13.0111 15.4129C13.0111 14.4661 13.6614 13.7923 14.5549 13.7923C15.3037 13.7923 15.7538 14.2314 16.0566 14.6602C16.2807 14.9745 16.4889 15.2089 16.6931 15.2089C16.9586 15.2089 17.0558 14.7236 17.0558 14.2575L17.0425 6.46595C17.0425 6.15386 17.007 6.09481 16.8303 6.0269C16.5413 5.91107 16.4312 5.73929 16.3988 5.36909L16.1449 2.73507C16.1289 2.48917 16.2421 2.22685 16.433 2.05197L17.5263 1.10515C17.6816 0.969331 17.8729 0.888101 18.0864 0.888101C18.3035 0.888101 18.4948 0.969331 18.6502 1.10515ZM16.9647 18.9769C17.2493 19.0674 17.401 19.3879 17.3069 19.6724L16.3512 22.4437C16.2504 22.7282 15.934 22.8893 15.6494 22.789C15.368 22.6954 15.2131 22.3749 15.3072 22.1032L16.2661 19.3155C16.37 19.0346 16.674 18.8734 16.9647 18.9769ZM19.9103 19.3155L20.8692 22.1032C20.9633 22.3873 20.8053 22.7171 20.5177 22.7921C20.2332 22.88 19.926 22.722 19.8284 22.4437L18.8726 19.6724C18.7785 19.3755 18.9365 19.0488 19.221 18.9738C19.5117 18.8859 19.8065 19.0408 19.9103 19.3155Z" fill-opacity="0.85"/><path d="M1.0623 25.7799C0.589134 26.8528 0.97886 28.1657 2.43337 28.1657C2.93137 28.1657 3.35836 27.9686 4.01883 27.4249L7.46038 24.5631C7.6051 24.4414 7.69034 24.405 7.76891 24.405C7.84393 24.405 7.92917 24.4414 8.07079 24.5631L11.519 27.4249C12.1759 27.9686 12.6065 28.1657 13.1045 28.1657C14.5621 28.1657 14.9518 26.8528 14.4755 25.7799L8.82887 13.0311L8.82887 7.02083C8.82887 4.13421 9.32026 1.97335 9.32026 1.55122C9.32026 1.15308 9.06504 0.888101 8.67356 0.888101C8.42456 0.888101 8.2044 1.02748 7.88259 1.3879L5.07809 4.60255C4.80956 4.91636 4.87081 5.31316 5.23255 5.47827C6.34538 5.98253 6.71561 6.13968 6.71561 7.02083L6.71561 13.0076ZM2.48126 26.2545L7.68233 14.3852L7.8555 14.3852L13.0566 26.2545C13.1778 26.543 12.9447 26.6837 12.7449 26.504L8.84088 23.2515C8.42807 22.9071 8.08407 22.738 7.76891 22.738C7.45686 22.738 7.10309 22.9071 6.69028 23.2515L2.796 26.504C2.58958 26.6837 2.35651 26.543 2.48126 26.2545ZM8.99149 20.3273C9.33194 20.3273 9.62046 20.049 9.62046 19.7085C9.62046 19.3583 9.3386 19.0729 8.99149 19.0729C8.6515 19.0729 8.35987 19.3717 8.35987 19.7085C8.35987 20.0423 8.65771 20.3273 8.99149 20.3273ZM9.68616 21.955C10.0333 21.955 10.3249 21.6763 10.3249 21.3327C10.3249 20.9825 10.0333 20.6909 9.68616 20.6909C9.34616 20.6909 9.04788 20.9989 9.04788 21.3327C9.04788 21.6763 9.34616 21.955 9.68616 21.955ZM10.3835 23.5792C10.727 23.5792 11.0289 23.2875 11.0289 22.9507C11.0289 22.6071 10.7306 22.3186 10.3835 22.3186C10.0533 22.3186 9.76163 22.6169 9.76163 22.9507C9.76163 23.2942 10.0435 23.5792 10.3835 23.5792Z" fill-opacity="0.85"/>',
    active: false,
  },
]);

const fetchAlbums = async () => {
  try {
    const api = getApiBase();
    const res = await fetch(`${api}/library`, {
      method: "GET",
      credentials: "include",
    });
    if (!res.ok) throw new Error("Not authorized");

    const data = await res.json();
    const list = data.albums || data;

    albums.value = list.map((a) => ({
      id: a.album_id ?? a.id,
      name: a.album_name ?? a.name,
      genre: a.album_genre ?? a.genre,
      artist: a.album_artist ?? a.artist,
      cover: a.album_cover ?? a.cover,
      date: a.album_date ?? a.date,
      songs: a.songs || a.Songs || [],

      date_added: a.date_added ?? a.uploaded_at ?? null,
      last_modified: a.last_modified ?? a.modified_at ?? null,
    }));

    username.value = data.username || "";

    console.log("[DEBUG] albums after mapping:", albums.value);
  } catch (err) {
    console.error(err);
    router.push("/");
  }
};

function parseSearch(query) {
  const tags = {
    artist: [],
    album: [],
    genre: [],
    song: [],
    text: [],
  };
  if (!query) return tags;
  query
    .split(/\s+/)
    .filter(Boolean)
    .forEach((token) => {
      const lower = token.toLowerCase();
      if (lower.startsWith("artist:")) tags.artist.push(lower.slice(7));
      else if (lower.startsWith("album:")) tags.album.push(lower.slice(6));
      else if (lower.startsWith("genre:")) tags.genre.push(lower.slice(6));
      else if (lower.startsWith("song:")) tags.song.push(lower.slice(5));
      else tags.text.push(lower);
    });
  return tags;
}

const filteredAlbums = computed(() => {
  if (!["Albums", "Recently Added"].includes(activeView.value))
    return albums.value;

  const tags = parseSearch(searchQuery.value);
  const hasTags =
    tags.artist.length ||
    tags.album.length ||
    tags.genre.length ||
    tags.text.length;
  if (!hasTags) return albums.value;

  const matches = (album) => {
    const name = (album.name || "").toLowerCase();
    const artist = (album.artist || "").toLowerCase();
    const genre = (album.genre || "").toLowerCase();
    if (tags.artist.length && !tags.artist.every((t) => artist.includes(t)))
      return false;
    if (tags.album.length && !tags.album.every((t) => name.includes(t)))
      return false;
    if (tags.genre.length && !tags.genre.every((t) => genre.includes(t)))
      return false;
    if (
      tags.text.length &&
      !tags.text.every(
        (t) => name.includes(t) || artist.includes(t) || genre.includes(t),
      )
    )
      return false;
    return true;
  };
  return albums.value.filter(matches);
});

const savedView = loadFromLocal("activeView", "Albums");
const activeView = ref(savedView || "Albums");

const setActive = (viewName) => {
  libraryNav.value.forEach((i) => (i.active = i.name === viewName));
  activeView.value = viewName;
};
const displayAlbums = computed(() => {
  const list = [...filteredAlbums.value];

  if (activeView.value === "Recently Added") {
    return list.sort(
      (a, b) => new Date(b.date_added || 0) - new Date(a.date_added || 0),
    );
  }

  if (activeView.value === "Albums") {
    return list.sort(
      (a, b) => new Date(b.last_modified || 0) - new Date(a.last_modified || 0),
    );
  }

  return list;
});

function showGuide(id) {
  if (!id) return;
  const key = `uploadGuideShown_${id}`;
  if (!localStorage.getItem(key)) {
    UploadGuide.value = true;
    setTimeout(() => {
      UploadGuide.value = false;
      localStorage.setItem(key, "true");
    }, 3500);
  }
}

function hideUploadGuide(userId) {
  const key = `uploadGuideShown_${userId}`;
  UploadGuide.value = false;
  localStorage.setItem(key, "true");
}

function triggerFilePicker(userId) {
  showGuide(userId);
  filepickerInput.value.click();
}

function sortBy(key) {
  if (sortKey.value === key) {
    sortOrder.value = sortOrder.value === "asc" ? "desc" : "asc";
  } else {
    sortKey.value = key;
    sortOrder.value = "asc";
  }
}

const uniqueArtists = computed(() => {
  const map = new Map();
  for (const album of albums.value) {
    const artistName = (album.artist || "Unknown").trim();
    if (!map.has(artistName)) {
      map.set(artistName, {
        name: artistName,
        albumCount: 0,
        covers: [],
      });
    }
    const entry = map.get(artistName);
    entry.albumCount = entry.albumCount + 1;
    if (album.cover) entry.covers.push(album.cover);
  }

  return [...map.values()].map((a) => ({
    name: a.name,
    albumCount: a.albumCount,
    randomCover: a.covers.length
      ? a.covers[Math.floor(Math.random() * a.covers.length)]
      : "/static/img/default-cover.jpg",
  }));
});

const selectArtist = (artistName) => {
  selectedArtist.value = artistName;
};

const artistAlbums = computed(() =>
  selectedArtist.value
    ? albums.value.filter((a) => a.artist === selectedArtist.value)
    : [],
);

const fuseArtists = computed(
  () =>
    new Fuse(uniqueArtists.value, {
      keys: ["name"],
      threshold: 0.4,
    }),
);

const filteredArtists = computed(() => {
  if (!searchQuery.value) return uniqueArtists.value;
  return fuseArtists.value.search(searchQuery.value).map((r) => r.item);
});

const groupedArtists = computed(() => {
  const groups = {};
  for (const artist of uniqueArtists.value) {
    const letter = artist.name[0]?.toUpperCase() || "#";
    if (!groups[letter]) groups[letter] = [];
    groups[letter].push(artist);
  }

  for (const letter in groups) {
    groups[letter].sort((a, b) => a.name.localeCompare(b.name));
  }

  return Object.fromEntries(
    Object.entries(groups).sort(([a], [b]) => a.localeCompare(b)),
  );
});

const fetchUploadStatus = async () => {
  try {
    const api = getApiBase();
    const res = await fetch(`${api}/upload/status`, {
      method: "GET",
      credentials: "include",
    });
    if (!res.ok) return;

    const data = await res.json();
    const progress = data.progress || [];

    const updated = uploadingSongs.value.map((song) => {
      const match = progress.find((p) => p.title === song.title);
      return match ? { ...song, ...match } : song;
    });

    const newSongs = progress.filter(
      (p) => !uploadingSongs.value.some((s) => s.title === p.title),
    );
    uploadingSongs.value = [...updated, ...newSongs];
    if (data.finished) {
      uploadTree.value = data.tree.replace(/^\s+/, "");
      console.log(uploadTree);
      uploadingSongs.value = uploadingSongs.value.filter((s) => s.done);
      stopProgressPolling();
      clearInterval(progressInterval);
    }
    if (
      uploadingSongs.value.length &&
      uploadingSongs.value.every((s) => s.done)
    ) {
      stopProgressPolling();
      clearInterval(progressInterval);

      await fetchAlbums();
    }
  } catch (err) {
    console.error("Error fetching upload status:", err);
  }
};

const startProgressPolling = () => {
  if (!progressInterval) {
    progressInterval = setInterval(fetchUploadStatus, 2000);
  }
};

const stopProgressPolling = () => {
  if (progressInterval) {
    clearInterval(progressInterval);
    progressInterval = null;
  }
};
const sortedUploading = computed(() =>
  [...uploadingSongs.value].sort((a, b) =>
    (a.artist || "").localeCompare(b.artist || ""),
  ),
);

const clearUploading = () => {
  uploadingSongs.value = [];
  stopProgressPolling();
};

const onDragOver = (e) => {
  if (e.dataTransfer?.types?.includes("Files")) {
    dragOver.value = true;
  }
};

const onDragEnter = (e) => {
  if (e.dataTransfer?.types?.includes("Files")) {
    dragOver.value = true;
  }
};

const onDragLeave = (e) => {
  dragOver.value = false;
};

const uploadFiles = async (files) => {
  const flacs = files.filter((f) =>
    f.name.toLowerCase().trim().endsWith(".flac"),
  );
  if (!flacs.length) {
    console.warn("No FLAC files found in selection/drop");
    return;
  }
  uploadingSongs.value = flacs.map((f) => ({
    artist: null,
    title: f.name.replace(/\.[^/.]+$/, ""),
    done: false,
  }));

  const formData = new FormData();
  flacs.forEach((f) => formData.append("files[]", f));

  startProgressPolling();
  const api = getApiBase();
  await fetch(`${api}/upload`, {
    method: "POST",
    credentials: "include",
    body: formData,
  });
};

const handleDrop = async (e) => {
  dragOver.value = false;
  const items = e.dataTransfer?.items;
  if (!items) return;

  const entries = [];
  for (const item of Array.from(items)) {
    const entry = item.webkitGetAsEntry?.();
    if (entry) entries.push(entry);
  }
  const files = [];
  const traverseEntry = async (entry, path = "") => {
    if (entry.isFile) {
      await new Promise((resolve) => {
        entry.file((file) => {
          if (file) {
            file.relativePath = path + file.name;
            files.push(file);
          }
          resolve();
        });
      });
    } else if (entry.isDirectory) {
      const reader = entry.createReader();
      const subEntries = await new Promise((res) => reader.readEntries(res));
      for (const subEntry of subEntries) {
        await traverseEntry(subEntry, path + entry.name + "/");
      }
    }
  };

  for (const entry of entries) {
    await traverseEntry(entry);
  }

  await uploadFiles(files);
};

const handleFiles = async (e) => {
  const list = e.target.files;
  if (!list || !list.length) return;
  const files = Array.from(list).map((f) => {
    const rel = f.webkitRelativePath || f.relativePath || f.name;
    f.relativePath = rel;
    return f;
  });
  await uploadFiles(files);
};

const groupedByGenre = computed(() => {
  const groups = {};
  for (const album of albums.value) {
    const genre = album.genre?.trim() || "Unknown";
    if (!groups[genre]) groups[genre] = [];
    groups[genre].push(album);
  }
  return Object.keys(groups)
    .sort((a, b) => a.localeCompare(b))
    .map((genre) => ({ genre, albums: groups[genre] }));
});
const uniqueGenres = computed(() => {
  const set = new Set();
  albums.value.forEach((a) => {
    if (a.genre) set.add(a.genre);
  });
  return Array.from(set).sort();
});

const groupedGenres = computed(() => {
  const groups = {};
  uniqueGenres.value.forEach((g) => {
    const letter = g[0].toUpperCase();
    if (!groups[letter]) groups[letter] = [];
    groups[letter].push(g);
  });
  return groups;
});

const lettersPresent = computed(() => Object.keys(groupedGenres.value).sort());

const scrollToLetter = (letter) => {
  const el = document.getElementById(letter);
  if (el) el.scrollIntoView({ behavior: "smooth", block: "start" });
};

const fuse = computed(() => {
  return new Fuse(albums.value, {
    keys: ["name", "artist"],
    threshold: 0.35,
    ignoreLocation: true,
  });
});

const filteredGenres = computed(() => {
  if (!searchQuery.value) return [];

  const tags = parseSearch(searchQuery.value);
  const results = albums.value.filter((album) => {
    const genre = (album.genre || "").toLowerCase();
    const name = (album.name || "").toLowerCase();
    const artist = (album.artist || "").toLowerCase();

    if (tags.genre.length && !tags.genre.every((t) => genre.includes(t)))
      return false;
    if (
      tags.text.length &&
      !tags.text.every(
        (t) => genre.includes(t) || name.includes(t) || artist.includes(t),
      )
    )
      return false;
    return true;
  });

  const groups = {};
  results.forEach((album) => {
    if (!groups[album.genre]) groups[album.genre] = [];
    groups[album.genre].push(album);
  });
  return Object.entries(groups).map(([genre, albums]) => ({ genre, albums }));
});

const filteredSongs = computed(() => {
  if (!searchQuery.value) return songs.value;

  const tags = parseSearch(searchQuery.value);
  const results = songs.value.filter((song) => {
    const title = (song.title || "").toLowerCase();
    const artist = (song.artist || "").toLowerCase();
    const album = (song.album || "").toLowerCase();
    const genre = (song.genre || "").toLowerCase();

    if (tags.song.length && !tags.song.every((t) => title.includes(t)))
      return false;
    if (tags.artist.length && !tags.artist.every((t) => artist.includes(t)))
      return false;
    if (tags.album.length && !tags.album.every((t) => album.includes(t)))
      return false;
    if (tags.genre.length && !tags.genre.every((t) => genre.includes(t)))
      return false;
    if (
      tags.text.length &&
      !tags.text.every(
        (t) =>
          title.includes(t) ||
          artist.includes(t) ||
          album.includes(t) ||
          genre.includes(t),
      )
    )
      return false;
    return true;
  });

  return results;
});

const displayTokens = computed(() => {
  const tokens = searchQuery.value.split(/\s+/).filter(Boolean);
  return tokens.map((raw) => {
    const m = raw.match(/^(album|artist|genre|song|text):(.*)$/i);
    if (!m) {
      return { raw, value: raw, isTag: false, class: "text-white" };
    }
    const kind = m[1].toLowerCase();
    const val = (m[2] || "").trim();
    const palette = {
      album: "bg-red-800/60 text-red-100 border border-red-500/60",
      artist: "bg-purple-800/60 text-purple-100 border border-purple-500/60",
      genre: "bg-green-800/60 text-green-100 border border-green-500/60",
      song: "bg-blue-800/60 text-blue-100 border border-blue-500/60",
      text: "bg-amber-800/60 text-amber-100 border border-amber-500/60",
    };
    const cls =
      palette[kind] || "bg-stone-700 text-stone-100 border border-stone-500/60";
    return {
      raw,
      label: kind,
      value: val,
      isTag: val.length > 0,
      class: cls,
    };
  });
});
const sortedSongs = computed(() => {
  const base = filteredSongs.value;
  return [...base].sort((a, b) => {
    const valA = (a[sortKey.value] || "").toString().toLowerCase();
    const valB = (b[sortKey.value] || "").toString().toLowerCase();
    if (valA < valB) return sortOrder.value === "asc" ? -1 : 1;
    if (valA > valB) return sortOrder.value === "asc" ? 1 : -1;
    return 0;
  });
});

function saveToLocal(key, value) {
  try {
    localStorage.setItem(key, JSON.stringify(value));
  } catch (e) {
    console.warn(e);
  }
}

function loadFromLocal(key, fallback) {
  try {
    const v = localStorage.getItem(key);
    return v ? JSON.parse(v) : fallback;
  } catch (e) {
    return fallback;
  }
}

const visibleColumns = ref(
  loadFromLocal("visibleColumns", [
    { key: "title", label: "Title" },
    { key: "artist", label: "Artist" },
    { key: "album", label: "Album" },
    { key: "genre", label: "Genre" },
    { key: "duration", label: "Duration" },
  ]),
);

watch(visibleColumns, (v) => saveToLocal("visibleColumns", v), { deep: true });
watch(activeView, (v) => saveToLocal("activeView", v));

let draggedIndex = null;

const dragIndex = ref(null);

function onDragStart(index) {
  dragIndex.value = index;
}

function onDrop(dropIndex) {
  if (dragIndex.value === null || dragIndex.value === dropIndex) return;
  const cols = [...visibleColumns.value];
  const [moved] = cols.splice(dragIndex.value, 1);
  cols.splice(dropIndex, 0, moved);
  visibleColumns.value = cols;
  dragIndex.value = null;
}

onMounted(async () => {
  window.addEventListener("resize", () => {
    windowWidth.value = window.innerWidth;
  });
  if (savedView && libraryNav.value.some((i) => i.name === savedView)) {
    setActive(savedView);
  }
  await fetchAlbums();
  const all = [];
  const api = getApiBase();
  const res = await fetch(`${api}/profile`, {
    credentials: "include",
  });
  if (res.ok) {
    const user = await res.json();
    userId.value = user.id;
  }
  for (const album of albums.value) {
    for (const s of album.songs || []) {
      all.push({
        id: s.song_id || s.id || `${album.id}-${s.title}`,
        title: s.title,
        duration: s.duration || 0,
        artist: s.artist || album.artist,
        album: album.name,
        genre: s.genre || album.genre || "Unknown",
        missing: !!s.missing,
      });
    }
  }
  songs.value = all;
  const viewFromQuery = route.query.view;
  if (
    typeof viewFromQuery === "string" &&
    libraryNav.value.some((i) => i.name === viewFromQuery)
  ) {
    setActive(viewFromQuery);
  } else {
    setActive(activeView.value);
  }
});
</script>

<style scoped>
.sidebar-slide-enter-active,
.sidebar-slide-leave-active {
  transition: all 0.3s ease;
}
.sidebar-slide-enter-from,
.sidebar-slide-leave-to {
  transform: translateY(-100%);
  opacity: 0;
}
.sidebar-slide-enter-to,
.sidebar-slide-leave-from {
  transform: translateY(0);
  opacity: 1;
}
</style>
