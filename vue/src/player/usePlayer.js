import { reactive, computed } from "vue";

const state = reactive({
  queue: [],
  history: [],
  currentIndex: -1,
  isPlaying: false,
  isShuffle: false,
  isRepeatOne: false,
  audioEl: null,
});

function _pickNextIndex() {
  if (!state.queue.length) return -1;

  if (state.isRepeatOne && state.currentIndex !== -1) {
    return state.currentIndex;
  }

  if (state.isShuffle) {
    if (state.queue.length === 1) return state.currentIndex;
    let next = state.currentIndex;
    while (next === state.currentIndex) {
      next = Math.floor(Math.random() * state.queue.length);
    }
    return next;
  }
  if (state.currentIndex + 1 < state.queue.length) {
    return state.currentIndex + 1;
  }

  return -1;
}

function _startAudioForCurrent() {
  if (!state.audioEl) return;
  const track = state.queue[state.currentIndex];
  if (!track) return;

  console.log("[player] start audio", {
    idx: state.currentIndex,
    id: track.id,
    src: track.streamUrl,
  });
  state.audioEl.src = track.streamUrl;
  state.audioEl
    .play()
    .then(() => {
      state.isPlaying = true;
    })
    .catch((err) => {
      console.error("[player] play() failed:", err);
      state.isPlaying = false;
    });
}

export function usePlayer() {
  const currentTrack = computed(() => {
    if (state.currentIndex < 0 || state.currentIndex >= state.queue.length)
      return null;
    return state.queue[state.currentIndex];
  });

  function addToHistory(track) {
    if (!track) return;
    state.history = state.history.filter((t) => t.id !== track.id);
    state.history.unshift(track);
    if (state.history.length > 200) state.history.pop();
  }

  function setAudioElement(el) {
    if (state.audioEl === el) return;
    state.audioEl = el || null;
    console.log("[player] setAudioElement", !!el);
    if (state.audioEl && state.currentIndex >= 0 && state.queue.length) {
      _startAudioForCurrent();
    }
  }

  function playNow(track, list) {
    const prev = currentTrack.value;
    console.log("[player] playNow", track?.id, "list size", list?.length);
    if (Array.isArray(list) && list.length) {
      state.queue = list.slice();
      const idx = state.queue.findIndex((t) => t.id === track.id);
      state.currentIndex = idx !== -1 ? idx : 0;
    } else {
      state.queue = [track];
      state.currentIndex = 0;
    }

    addToHistory(prev);

    _startAudioForCurrent();
  }

  function enqueueTrack(track) {
    state.queue.push(track);
  }

  function enqueueAlbum(trackList) {
    if (!Array.isArray(trackList) || !trackList.length) return;
    state.queue.push(...trackList);
  }

  function togglePlay() {
    if (!state.audioEl) return;
    if (state.isPlaying) {
      state.audioEl.pause();
      state.isPlaying = false;
    } else {
      if (!currentTrack.value && state.queue.length) {
        state.currentIndex = 0;
        _startAudioForCurrent();
      } else {
        state.audioEl
          .play()
          .then(() => (state.isPlaying = true))
          .catch((err) => console.error("[player] resume failed", err));
      }
    }
  }

  function playNext() {
    if (!state.queue.length) return;
    const nextIdx = _pickNextIndex();
    if (nextIdx === -1) {
      state.isPlaying = false;
      return;
    }

    const prev = currentTrack.value;
    state.currentIndex = nextIdx;
    addToHistory(prev);
    _startAudioForCurrent();
  }

  function playPrev() {
    if (!state.queue.length) return;
    if (state.history.length) {
      const prev = state.history.shift();
      const idx = state.queue.findIndex((t) => t.id === prev.id);
      if (idx !== -1) {
        state.currentIndex = idx;
        _startAudioForCurrent();
        return;
      }
    }
    // якщо історії нема — просто попередній по індексу
    if (state.currentIndex > 0) {
      state.currentIndex -= 1;
      _startAudioForCurrent();
    }
  }

  function onTrackEnded() {
    playNext();
  }

  function toggleShuffle() {
    state.isShuffle = !state.isShuffle;
  }

  function toggleRepeatOne() {
    state.isRepeatOne = !state.isRepeatOne;
  }

  function clearQueue() {
    state.queue = [];
    state.currentIndex = -1;
    state.isPlaying = false;
  }

  return {
    state,
    currentTrack,
    setAudioElement,
    playNow,
    enqueueTrack,
    enqueueAlbum,
    togglePlay,
    playNext,
    playPrev,
    toggleShuffle,
    toggleRepeatOne,
    clearQueue,
    onTrackEnded,
  };
}
