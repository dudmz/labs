import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
    vus: 10,
    duration: '30s',
};

export default () => {
    sleep(3)
    http.get("http://damazios-MacBook-Pro.local:8080/messages")
}
