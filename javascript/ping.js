const { exec } = require('child_process');

const host = process.argv[2];
if (!host) {
    console.error('Usage: node ping.js <host>');
    process.exit(1);
}

exec(`ping -c 4 ${host}`, (error, stdout) => {
    if (error) {
        console.error(`Error: ${error.message}`);
        return;
    }
    console.log(stdout);
});
