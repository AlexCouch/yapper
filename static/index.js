import WaveSurfer from 'https://unpkg.com/wavesurfer.js@7/dist/wavesurfer.esm.js'
const container = document.getElementById('waveformContainer')
const rect = container.getBoundingClientRect()
const width = rect.width
const height = rect.height

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
})

wavesurfer.on('interaction', ()=>{
    wavesurfer.play()
})
