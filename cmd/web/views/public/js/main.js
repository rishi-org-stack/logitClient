window.onload = ()=>{
	let ok = document.getElementById("ok");
	let sIdea = document.getElementById("s-idea");
	let sEvent = document.getElementById("s-event");
	let sNews = document.getElementById("s-news");
	let sBlog = document.getElementById("s-blog");
	let path = document.querySelector("#path");
    let c=0;

    // path.onclick=()=>{
    //     console.log("path clicked");
    // }
	ok.onclick=()=>{
        console.log("ok clicked")
    }
    ok.onclick=()=>{

        console.log("ok clicked")
    }
    sIdea.onclick=()=>{
        c++;
        sIdea.style.padding="10px"
        sIdea.style.paddingLeft="30px"
        sIdea.style.paddingRight="30px"
        sIdea.style.borderRadius="5px"
        // sIdea.style.width=150
        sIdea.style.backgroundColor="brown"
        
        console.log("ok clicked")
        if (c==1){
            let idea = document.createElement("span")
            idea.innerHTML="idea"
            path.appendChild(idea)
            idea.onclick=()=>{
                console.log("idea");
            }
        }
        
        sNews.style.backgroundColor="transparent"
        sBlog.style.backgroundColor="transparent"
        sEvent.style.backgroundColor="transparent"
    }
    sIdea.onmouseleave=()=>{
        // c=0;
    }
    sEvent.onclick=()=>{
        sEvent.style.backgroundColor="brown"
        sEvent.style.padding="10px"
        sEvent.style.paddingLeft="30px"
        sEvent.style.paddingRight="30px"
        sEvent.style.borderRadius="5px"
        console.log("ok clicked")

        sBlog.style.backgroundColor="transparent"
        sIdea.style.backgroundColor="transparent"
        sNews.style.backgroundColor="transparent"
    };
    sNews.onclick=()=>{
        sNews.style.backgroundColor="brown"
        sNews.style.padding="10px"
        sNews.style.paddingLeft="30px"
        sNews.style.paddingRight="30px"
        sNews.style.borderRadius="5px"
        console.log("ok clicked")


        sBlog.style.backgroundColor="transparent"
        sIdea.style.backgroundColor="transparent"
        sEvent.style.backgroundColor="transparent"
        // sNews.style.backgroundColor="transparent"
    }
    sBlog.onclick=()=>{
        sBlog.style.backgroundColor="brown"
        sIdea.style.backgroundColor="transparent"
        sEvent.style.backgroundColor="transparent"
        sNews.style.backgroundColor="transparent"
        sBlog.style.padding="10px"
        sBlog.style.paddingLeft="30px"
        sBlog.style.paddingRight="30px"
        sBlog.style.borderRadius="5px"
        console.log("ok clicked")
    }


    function press(){
        console.log("ok pressed");
    }
}