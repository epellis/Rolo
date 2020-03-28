import { writable } from 'svelte/store';

class UserStore {
    static defaultUser() {
        return { isLoggedIn: false };
    }

    constructor() {
        let storedUserData;
        try {
            storedUserData = JSON.parse(localStorage.getItem("userData") || "");
            console.log("Loading User:", storedUserData);
        } catch {
            storedUserData = UserStore.defaultUser();
            console.log("Loading Default User:", storedUserData);
        }
        const { subscribe, set, update } = writable(storedUserData);
        this.subscribe = subscribe
        this.set = set
        this.update = update
    }

    subscribe() {
        return this.subscribe;
    }

    logIn(userData) {
        userData.isLoggedIn = true;
        localStorage.setItem("userData", JSON.stringify(userData));
        this.set(userData)
    }

    logOut() {
        console.log("Logging Out");
        localStorage.setItem("userData", JSON.stringify(UserStore.defaultUser()));
        this.set(UserStore.defaultUser());
    }
}

export const userStore = new UserStore();
