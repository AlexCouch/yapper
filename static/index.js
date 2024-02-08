import WaveSurfer from 'https://unpkg.com/wavesurfer.js@7/dist/wavesurfer.esm.js'

const container = document.getElementById('waveformContainer')
const rect = container.getBoundingClientRect()
const width = rect.width
const height = rect.height

const play_fa = 'fa-play'
const pause_fa = 'fa-pause'

const wavesurfer = WaveSurfer.create({
    container: "#waveform",
    waveColor: "rgb(3 105 161)",
    progressColor: "rgb(125 211 252)",
    cursorWidth: 0,
    url: "/audio.mp3",
    width: width,
    height: height,
    barWidth: 9,
    barHeight: 1.1,
    barRadius: 12,
    barGap: 2,
    dragToSeek: true,
})

function setPauseButton(){
    let yap_play = $('#yap_play')
    yap_play.removeClass(play_fa)
    yap_play.addClass(pause_fa)
}

function setPlayButton(){
    let yap_play = $('#yap_play')
    yap_play.removeClass(pause_fa)
    yap_play.addClass(play_fa)
}

function stop(){
    wavesurfer.pause()
    wavesurfer.seekTo(0)
    setPlayButton()
}

function pause(){
    console.log("Pausing...")
    wavesurfer.pause()
    setPlayButton()
}

function play(){
    console.log("Playing...")
    wavesurfer.play()
    setPauseButton()
}

$(document).on('click', function(e){
    if(!$(event.target).closest('#yap').length){
        stop()
    }
})

wavesurfer.on('ready', function(){
    setPlayButton()
})

wavesurfer.on('finish', function(){
    stop()
})

$('#play_button').on('click', function(e){
    if(wavesurfer.isPlaying()){
        pause()
    }else{
        play()
    }
})
