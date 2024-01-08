import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {  faPlus } from '@fortawesome/free-solid-svg-icons';

export const AddFileBtn = () => {
    return(
        <button 
        className="p-3 rounded-md bg-indigo-700 text-white font-semibold w-full mt-5 hover:opacity-75"
        >
           <FontAwesomeIcon icon={faPlus} /> New
        </button>
    )
}