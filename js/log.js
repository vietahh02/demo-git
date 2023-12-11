import {log, error, warn} from "../js/constants.js"

function logger (mess, type = log) {
    console[type](mess)
}

export default logger