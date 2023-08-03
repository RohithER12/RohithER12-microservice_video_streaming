window.addEventListener('load', function() {
    fetch('/video/all') 
    .then(async response => {
        let res = await response.json();
        console.log(res);

        res.Video.forEach(video => { 
            console.log(video);
            const videoElement = document.createElement('video');
            videoElement.controls = true;
            
            videoElement.autoplay = true;       
            
            if (Hls.isSupported()) {
                const hls = new Hls();
                hls.loadSource('/stream/' + video.VideoId + '/playlist.m3u8');
                hls.attachMedia(videoElement);
            } else if (videoElement.canPlayType('application/vnd.apple.mpegurl')) {
                videoElement.src = '/stream/' + video.VideoId + '/playlist.m3u8';
            } else {
                console.error('HLS is not supported on this browser.');
            }
            
            document.getElementById('videoContainer').appendChild(videoElement);
        });
    })
    .catch(error => {
        console.error('Error fetching video data:', error);
    });
});






// window.addEventListener('load', function() {
// fetch('/video/all') 
// .then(response => response.json())
// .then(data => {
//     data.Video.forEach(video => { 
//         console.log(video);
//         const videoElement = document.createElement('video');
//         videoElement.controls = true;
//         const sourceElement = document.createElement('source');

//         sourceElement.src = '/stream/' + video.VideoId + '/playlist.m3u8';
//         sourceElement.type = 'video/mp4'; 
//         videoElement.appendChild(sourceElement);
//         document.getElementById('videoContainer').appendChild(videoElement);
//     });
// })
// .catch(error => {
//     console.error('Error fetching video data:', error);
// });
// })