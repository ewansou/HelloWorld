# GoSchool - Microservices (Activity #3: Simple Web Calculator with GitHub)

Go Microservices 2 Assessment

Build and Test using YML and Github Actions

Collaborators: @gabrielleeyj 

GitHub Actions

(1) mygo.yml - On PUSH and PULL_REQUEST on main branch, script will run, build and do tests (using Goblin). Image scaling has also been automated too. Project images in the folder 'project-images' are being scaled down by the github action (https://github.com/marketplace/actions/scale-images). New images are then subsequently automatically push up back to repo

(2) mobileNotification.yml - I will receive a SMS notication when changes has been made to the main branch of this repo. This action is done using Twilio SMS github action (https://github.com/marketplace/actions/twilio-sms).

(3) release.yml - To automate release/deployment process.