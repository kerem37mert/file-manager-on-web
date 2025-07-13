const buttons = document.querySelector(".buttons");
const newFile = document.getElementById("newFile");
const newFolder = document.getElementById("newFolder");
const newForm = document.querySelector(".newForm");
const closeFormBtn = document.getElementById("closeFormBtn");

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
    openForm("formNewFile");
});

newForm.addEventListener("submit", (event) => {
    event.preventDefault();

    const name = document.querySelector(".newForm input").value;

    if(name == "")
        return;
    console.log(name);
});