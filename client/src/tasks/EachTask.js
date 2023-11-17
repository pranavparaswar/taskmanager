import { Link } from "react-router-dom"
import { L } from "../config"
import axios from "axios"
import moment from "moment"
import { useState } from 'react'
import { format } from 'date-fns'
import { API_URL } from '../config';

export default function EachTask({task, fetchData}) {
    const [nameValue, setNameValue] = useState(task.name)
    const [noteValue, setNoteValue] = useState(task.note)
    const [dateValue, setdateValue] = useState(moment(task.date).format('YYYY-MM-DD'))
    const [isCompleted, setIsCompleted] = useState(false);
    const formattedDate = format(new Date(task.date), 'dd/MM/yyyy');

    const toggleCompleted = () =>{
        setIsCompleted(!isCompleted);
    };

    const openModal = () => {
        document.getElementById('new-modal-' + task.id).classList.remove("hidden");
    }

    const closeModal = () => {
        document.getElementById('new-modal-' + task.id).classList.add("hidden");
    }

    const completeForm = () => {
        closeModal()
        fetchData()
    }

    const updateTask = (e) => {
    e.preventDefault();
    var form = document.getElementById(`editform-${task.id}`);
    var formData = new FormData(form);
    
    console.log(task.id);
    console.log(formData);

    axios.patch(`${API_URL}/task/${task.id}`, formData)
        .then(res => {
            completeForm();
        })
        .catch((error) => {
            if (error.response && error.response.status === 500 && error.response.data.includes('duplicate key value violates unique constraint "unique_taskname"')) {
                alert('Task name must be unique. Please choose a different name.');
            } else {
                console.log(error.response);
            }
        });
};


    const deleteTask = async () =>{
        console.log("Task ID to be deleted:", task.id);
        if (window.confirm("Do you want to delete this task")==true){
            await axios.delete(`${API_URL}/task/${task.id}`)
            .then(res =>fetchData())
            .catch(error => console.log(error.response))
        } else{
            console.log("Cancelled")
        }
    }

    return (
        <div className="bg-slate-1000 rounded-lg mb-4 p-4 hover:border-purple-700">
            <div>
                <div>
                    <div onClick={toggleCompleted} className={`cursor-pointer font-medium ${isCompleted ? 'line-through' : ''}`}>
                        {task.name}
                    </div>
                    <div className="text-slate-400">Note : {task.note} </div>
                    <div className="text-slate-400">Due By {formattedDate} </div>
                </div>
                <div className="text-sm flex space-x-4 mt-4">
                    <button onClick={toggleCompleted}>Mark Complete</button>
                    <button onClick={openModal}>Edit Task</button>
                    <button onClick={deleteTask} className="text-red-600">Delete Task</button>
                </div>
            </div>

            {/*Start Modal */}
            <div className="relative z-10 hidden" aria-labelledby="modal-title" role="dialog" aria-modal="true" id={`new-modal-${task.id}`}>
                <div className="fixed inset-0 bg-black bg-opacity-70 transition-opacity"></div>

                <div className="fixed z-10 inset-0 overflow-y-auto">
                    <div className="flex items-end justify-center min-h-sreen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
                        <span className="hidden sm:inline-nock sm:align middle sm:h-screen" aria-hidden="true">&#8203;</span>

                        <div className="relative inline-block align-middle bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:max-w-lg w-full">
                            <form id={`editform-${task.id}`} onSubmit={updateTask} method="post">
                                <div className="bg-white">
                                    <div className="flex justify-between px-8 py-4 border-b">
                                        <h1 className="font-medium">Update Task</h1>
                                        <button type="button" onClick={closeModal}>Close</button>
                                    </div>
                                    <div className="px-8 py-8">
                                        <div className="mb-5">
                                            <label className="block text-gray-700 text-sm font-bold mb-2">Task Name</label>
                                            <input type="text" name="TaskName" value={nameValue} onChange={(e) => setNameValue(e.target.value)} className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" required />
                                        </div>
                                        <div className="mb-5">
                                            <label className="block text-gray-700 text-sm font-bold mb-2">Note</label>
                                            <input type="text" name="note" value={noteValue} onChange={(e) => setNoteValue(e.target.value)} className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" required />
                                        </div>
                                        <div className="mb-5">
                                            <label className="block text-gray-700 text-sm font-bold mb-2">Date</label>
                                            <input type="date" name="deadline" value={dateValue} onChange={(e) => setdateValue(e.target.value)} className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" required />
                                        </div>
                                         <div className="flex justify-end">
                                            <button className="bg-blue-500 text-white py-1.5 px-4 rounded" type="submit">Submit</button>
                                        </div>
                                    </div>
                                </div>
                            </form>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )

}