import WaveSurfer from 'https://unpkg.com/wavesurfer.js@7/dist/wavesurfer.esm.js'

var wavesurfers = []
let play_fa = 'fa-play'
let pause_fa = 'fa-pause'

export function render(id, selector, url){
    const yapContainer = `.yap${id}`
    const playButton = `.yap_play${id}`

    let outerContainer = $(selector)[id]
    console.log(outerContainer)
    let rect = outerContainer.getBoundingClientRect()
    console.log(rect)
    let width = rect.width
    let height = rect.height

    let trackid = "t" + id
    let container = $("<div id='" + trackid + "'>").appendTo(outerContainer)

    

    wavesurfers[trackid] = WaveSurfer.create({
        container: `#${trackid}`,
        waveColor: "rgb(3 105 161)",
        progressColor: "rgb(125 211 252)",
        cursorWidth: 0,
        url: url,
        width: width,
        height: height,
        barWidth: 9,
        barHeight: 1.1,
        barRadius: 12,
        barGap: 2,
        dragToSeek: true,
    })

    let wavesurfer = wavesurfers[trackid]

    $(document).on('click', function(e){
        if(!$(event.target).closest(yapContainer).length){
            stop(wavesurfer, playButton)
        }
    })

    wavesurfer.on('ready', function(){
        setPlayButton(wavesurfer, playButton)
    })

    wavesurfer.on('finish', function(){
        stop(wavesurfer, playButton)
    })

    $(playButton).on('click', function(e){
        if(wavesurfer.isPlaying()){
            pause(wavesurfer, playButton)
        }else{
            play(wavesurfer, playButton)
        }
    })
}

function setPauseButton(wavesurfer, playButton){
    let yap_play = $(playButton)
    yap_play.removeClass(play_fa)
    yap_play.addClass(pause_fa)
}

function setPlayButton(wavesurfer, playButton){
    let yap_play = $(playButton)
    console.log(yap_play)
    if(yap_play.hasClass(pause_fa)){
        yap_play.removeClass(pause_fa)
    }
    yap_play.addClass(play_fa)
}

function stop(wavesurfer, playButton){
    wavesurfer.pause()
    wavesurfer.seekTo(0)
    setPlayButton(wavesurfer, playButton)
}

function pause(wavesurfer, playButton){
    console.log("Pausing...")
    wavesurfer.pause()
    setPlayButton(wavesurfer, playButton)
}

function play(wavesurfer, playButton){
    console.log("Playing...")
    wavesurfer.play()
    setPauseButton(wavesurfer, playButton)
}

