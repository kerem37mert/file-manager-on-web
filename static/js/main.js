const buttons = document.querySelector(".buttons");
const newFile = document.getElementById("newFile");
const newFolder = document.getElementById("newFolder");
const newForm = document.querySelector(".newForm");
const closeFormBtn = document.getElementById("closeFormBtn");
const error = document.querySelector(".error");

const openForm = (id) => {
    newForm.id = id;
    buttons.style.display = "none";
    newForm.style.display = "flex";
}

const closeForm = () => {
    newForm.style.display = "none";
    buttons.style.display = "flex";
}

closeFormBtn.addEventListener("click", closeForm);

newFile.addEventListener("click", () => {
    openForm("formNewFile");
});

newFolder.addEventListener("click", () => {
    openForm("formNewFolder");
});

newForm.addEventListener("submit", (event) => {
    event.preventDefault();

    const name = document.querySelector(".newForm input").value;

    if(name == "")
        return;
    
    let isFolder = true;
    if(newForm.id == "formNewFile")
        isFolder = false;

    const path = window.location.pathname;
    const pathName = path == "/" ? `/${name}` : `${path}/${name}`; 
    
    fetch(`/filemanagerapi/new?name=${pathName}&isFolder=${isFolder}`, {
        method: "GET"
    })
    .then(response => response.json())
    .then(data => {
        if(data != null) {
            error.style.display = "block";
            error.innerHTML = data.message;
            return
        }
        window.location.reload();
    })
    .catch(err => {
        console.log(err);
    })
});