const sidebar = document.querySelector(".Sidebar_root");
const sidebar_container = document.querySelector(".Sidebar_container");

function OpenSidebar() {
    sidebar.style.opacity = 1;
    sidebar.style.zIndex = 750;
    sidebar_container.style.transform = "translate(0)";
}

function CloseSidebar() {
    sidebar.style.opacity = 0;
    sidebar.style.zIndex = -1;
    sidebar_container.style.transform = "translate(25rem)";
}

const sb_accordions = document.querySelectorAll('.Sidebar_accordion');
sb_accordions.forEach((accordion) => {
  accordion.addEventListener('click', ToggleSidebarSublist.bind(accordion))
});

function ToggleSidebarSublist() {
  // Make Sidebar_item text BOLD
  this.classList.toggle("Sidebar_item_selected");

  // Toggle Chevron direction
  const imgs = this.querySelectorAll("img");
  for (let i = 0; i < imgs.length; i++) {
    if (imgs[i].style.opacity == "") {
      sublist[i].style.opacity = 1;
    } else {
      sublist[i].style.opacity = "";
    }
      // imgs[i].classList.toggle('hidden');
  }

  // Toggle Sublist
  const sublist = this.parentElement.querySelectorAll(".Sidebar_subitem");
  for (let i = 0; i < sublist.length; i++) {
    if (sublist[i].style.opacity == "") {
      sublist[i].style.height = "3rem";
      sublist[i].style.opacity = 1;
    } else {
      sublist[i].style.height = 0;
      sublist[i].style.opacity = "";
    }
  }
}
