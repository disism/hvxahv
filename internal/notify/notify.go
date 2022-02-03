package notify

import (
	"gorm.io/gorm"
)

// Submit the client-registered data to hvxahv, and the pushed endpoint,
// p256dh and auth. hvxahv will send user notification message through the address.
// https://datatracker.ietf.org/doc/html/rfc6108

// https://developer.mozilla.org/en-US/docs/Web/API/Push_API
// {
//   "subscription":{
//     "endpoint":"https://sg2p.notify.windows.com/w/?token=BQYAAACLy9DPk5%2bkrOfxjm4cTtROW5peMRUg1m609l8lvgOje%2bzvQXtv2zW9pg286gAo8X67OlhkvCN7MC90MR6YkdBLk4aXHmiXd941QtkMLoDIv8ep2kXLgfKs7VEhJE1%2bXA9XOeoED%2brmFlojNhkxO%2b9N93cjob54Jo0nzfXUDUyOSUnXkgfcbnx7a0M4u9dExBCXIpfvIaUIvBmhqGhWgyW0KrgUv%2bj7R73SErOWdoYfOOISybv7Io55NAbrjRFRMIFzkepB3LLd2F2KZKQTg3o2f4nghRcM0qRqkdVRregsgl6eAFyzndhxVTff%2fBSiTP0%3d",
//     "expirationTime":null,
//     "keys":{
//       "p256dh":"BM9H6kiGNliWfhI23CrawUefVwsCYkIFZCsggtNTYSNy4Y5BzEJVrK3iM_0ZMzRndKYj2z7fXmBzxoQSrXnvxsQ",
//       "auth":"EJOFwKDlaVpDUV7uxcNPwg"
//     }
//   }
// }

type Notifies struct {
	gorm.Model

	DeviceID uint   `gorm:"primaryKey;type:bigint;device_id"`
	Endpoint string `gorm:"type:text;endpoint"`
	P256dh   string `gorm:"type:text;p256dh"`
	Auth     string `gorm:"type:text;auth"`
}
