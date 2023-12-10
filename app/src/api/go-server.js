import axios from "axios";

class GoServer {
  constructor() {
    this.client = axios.create({
      baseURL: "http://localhost:8080",
      timeout: 1000,
      headers: {
        "Content-Type": "application/json",
      },
    });
  }

  async getUsers() {
    try {
      const response = await this.client.get("/users/");
      return response;
    } catch (error) {
      console.error("Erreur lors de la récupération des utilisateurs :", error.message);
      throw error;
    }
  }

  async signUp(username, email, password) {
    try {
      const body = {
        username: username,
        email: email,
        password: password,
      };

      const response = await this.client.post("/auth/signup", body);
      return response;
    } catch (error) {
      console.error("Erreur lors de l'inscription :", error.message);
      throw error;
    }
  }

  async logIn(username, password) {
    try {
      const body = {
        username: username,
        password: password,
      };

      const response = await this.client.post("/auth/login", body);
      return response;
    } catch (error) {
      console.error("Erreur lors de la connexion :", error.message);
      throw error;
    }
  }

  async createNote(user_id, title, content) {
    try {
      const body = {
        user_id: parseInt(user_id),
        title: title,
        content: content,
      };

      const response = await this.client.post("/notes/", body);
      return response;
    } catch (error) {
      console.error("Erreur lors de la création de la note :", error.message);
      throw error;
    }
  }

  async updateNoteForUser(user_id, note_id, title, content) {
    try {
      const body = {
        user_id: user_id,
        note_id: note_id,
        title: title,
        content: content,
      };

      const response = await this.client.put("/notes/", body);
      return response;
    } catch (error) {
      console.error("Erreur lors de la mise à jour de la note :", error.message);
      throw error;
    }
  }

  async getNotesForUser(user_id) {
    try {
      const body = {
        user_id: parseInt(user_id),
      };

      const response = await this.client.get("/notes/", { params: body });
      return response;
    } catch (error) {
      console.error("Erreur lors de la récupération des notes :", error.message);
      throw error;
    }
  }

  async deleteNoteForUser(user_id, note_id) {
    try {
      const body = {
        user_id: user_id,
        note_id: note_id,
      };

      const response = await this.client.delete("/notes/", { data: body });
      return response;
    } catch (error) {
      console.error("Erreur lors de la suppression de la note :", error.message);
      throw error;
    }
  }
}

const goServer = new GoServer();
export default goServer;
