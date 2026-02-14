<script lang="ts">
    import { onMount } from 'svelte';
    import { page } from '$app/stores';

    let { videoId = "dQw4w9WgXcQ" } = $props();

    let player: any = $state(null);
    let currentTime = $state(0);
    let duration = $state(0);
    let isPaused = $state(true);

    // Calculate progress for the timeline
    const progress = $derived(duration > 0 ? (currentTime / duration) * 100 : 0);

    onMount(() => {
        // Load YouTube API
        const tag = document.createElement('script');
        tag.src = "https://www.youtube.com/iframe_api";
        document.head.appendChild(tag);

        (window as any).onYouTubeIframeAPIReady = () => {
            player = new (window as any).YT.Player('yt-player', {
                height: '100%',
                width: '100%',
                videoId,
                playerVars: { 'controls': 0, 'disablekb': 1, 'modestbranding': 1 },
                events: {
                    'onReady': (e: any) => duration = e.target.getDuration(),
                    'onStateChange': (e: any) => {
                        isPaused = e.data !== 1; // 1 is PLAYING
                    }
                }
            });
        };

        // Update timeline 10 times per second
        const interval = setInterval(() => {
            if (player && !isPaused) {
                currentTime = player.getCurrentTime();
            }
        }, 100);

        return () => clearInterval(interval);
    });

    const togglePlay = () => {
        if (isPaused) player.playVideo();
        else player.pauseVideo();
    };
</script>

<div class="video-container">
    <div id="yt-player"></div>
    
    <div class="controls">
        <button onclick={togglePlay}>
            {isPaused ? '▶' : '⏸'}
        </button>
        
        <div class="timeline-wrapper">
            <div class="timeline-bg">
                <div class="timeline-fill" style:width="{progress}%"></div>
            </div>
            <span class="time-label">{Math.floor(currentTime)}s / {Math.floor(duration)}s</span>
        </div>
    </div>
</div>

<style>
    .video-container { position: relative; width: 100%; aspect-ratio: 16/9; background: #000; overflow: hidden; border-radius: 12px; }
    .controls { position: absolute; bottom: 0; width: 100%; background: linear-gradient(transparent, rgba(0,0,0,0.8)); padding: 20px; display: flex; align-items: center; gap: 15px; }
    button { background: #fff; border: none; border-radius: 50%; width: 40px; height: 40px; cursor: pointer; font-size: 1.2rem; }
    .timeline-wrapper { flex-grow: 1; position: relative; }
    .timeline-bg { background: rgba(255,255,255,0.2); height: 6px; border-radius: 3px; overflow: hidden; }
    .timeline-fill { background: #ff0000; height: 100%; transition: width 0.1s linear; }
    .time-label { color: #fff; font-size: 0.8rem; font-family: monospace; }
</style>