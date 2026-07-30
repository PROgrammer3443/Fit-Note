var navLinks = document.querySelectorAll("nav a");
for (var i = 0; i < navLinks.length; i++) {
	var link = navLinks[i]
	if (link.getAttribute('href') == window.location.pathname) {
		link.classList.add("live");
		break;
	}
}

const mealModal = document.getElementById("mealModal");
const openMealModal = document.getElementById("openMealModal");
const closeMealModal = document.getElementById("closeMealModal");

if (mealModal && openMealModal && closeMealModal) {
    openMealModal.addEventListener("click", () => {
        mealModal.classList.add("active");
    });

    closeMealModal.addEventListener("click", () => {
        mealModal.classList.remove("active");
    });

    mealModal.addEventListener("click", (e) => {
        if (e.target === mealModal) {
            mealModal.classList.remove("active");
        }
    });
}

document.addEventListener("DOMContentLoaded", () => {
    const modal = document.getElementById("mealModal");
    const openBtn = document.getElementById("openMealModal");
    const closeBtn = document.getElementById("closeMealModal");

    if (modal && openBtn && closeBtn) {
        openBtn.addEventListener("click", () => {
            modal.classList.add("active");
        });

        closeBtn.addEventListener("click", () => {
            modal.classList.remove("active");
        });

        modal.addEventListener("click", (e) => {
            if (e.target === modal) {
                modal.classList.remove("active");
            }
        });
    }
});

function toggleEntries(id, btn) {
    const row = document.getElementById(id);

    if (!row) {
        return;
    }

    row.classList.toggle("active");

    const icon = btn.querySelector(".material-symbols-outlined");

    if (!icon) {
        return;
    }

    if (row.classList.contains("active")) {
        icon.textContent = "arrow_drop_up";
    } else {
        icon.textContent = "arrow_drop_down";
    }
}

document.querySelectorAll(".add-food-btn").forEach(btn => {
    btn.addEventListener("click", () => {
        openFoodModal(btn.dataset.mealId);
    });
});

function openFoodModal(mealID) {
    document.getElementById("mealIDInput").value = mealID;

    document
        .getElementById("foodModal")
        .classList.add("active");
}

function closeFoodModal() {
    document
        .getElementById("closeFoodModal")
        .classList.remove("active");
}

function openSwapModal(id) {
    const el = document.getElementById("currentExerciseID");
    el.value = id;

    console.log("Current Exercise ID:", el.value);

    document.getElementById("swapModal").classList.add("active");
}

function closeSwapModal() {

    document.getElementById("swapModal").classList.remove("active");
}


document.addEventListener("DOMContentLoaded", () => {

    const modal = document.getElementById("swapModal");
    const closeBtn = document.getElementById("closeSwapModal");
    const currentInput =
        document.getElementById("currentExerciseInput");

    document.querySelectorAll(".swap-btn").forEach(btn => {

        btn.addEventListener("click", () => {

            currentInput.value =
                btn.dataset.exercise;

            modal.classList.add("active");
        });
    });

    closeBtn.addEventListener("click", () => {

        modal.classList.remove("active");
    });

    modal.addEventListener("click", e => {

        if (e.target === modal) {
            modal.classList.remove("active");
        }
    });
});
