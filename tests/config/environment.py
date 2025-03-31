API_URI = "/api/v1/pokemon"

class Environment:
    def __init__(self, base_url, port):
        self.base_url = base_url
        self.port = port

    def get_base_uri(self):
        return f"http://{self.base_url}:{self.port}"
    
    def get_api_base_uri(self):
        return f"{self.get_base_uri()}{API_URI}"
