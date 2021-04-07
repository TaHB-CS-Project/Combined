$(document).ready(function(){
    var maxField = 10; 
    var add_diagnosis_button = $('.add_diagnosis_button'); 
    var diagnosis_wrapper = $('.diagnosis_wrapper'); 
    var diagnosisHTML = '<div style="display:flex;"><select name="diagnosis[]" class="form-control" id="diagnosis" style="margin-top:10px;"><option>Choose one</option><option>-----------</option><option>Intubation</option><option>UVC</option><option>UAC</option><option>PICC</option><option>Peritoneal tap</option><option>Urinary catheterization</option><option>Lumbar puncture</option><option>Ventricular tap</option><option>PIV</option><option>Chest tube placement</option></select><a href="javascript:void(0);" class="remove_diagnosis_button" title="Remove field"><i class="fas fa-minus-circle" style="font-size:22px; margin:15px 0 0 5px;"></i></a></div>'; 
    var x = 1; //Initial field counter is 1
    

    $(add_diagnosis_button).click(function(){
        if(x < maxField){ 
            x++; 
            $(diagnosis_wrapper).append(diagnosisHTML);
        }
    });
    

    $(diagnosis_wrapper).on('click', '.remove_diagnosis_button', function(e){
        e.preventDefault();
        $(this).parent('div').remove(); 
        x--; 
    });
});

$(document).ready(function(){
    var maxField = 10;
    var add_procedure_button = $('.add_procedure_button');
    var procedure_wrapper = $('.procedure_wrapper'); 
    var procedureHTML = '<div style="display:flex;"><select name="procedure[]" class="form-control" id="procedure" style="margin-top:10px;"><option>Choose one</option><option>-----------</option><option>Intubation</option><option>UVC</option><option>UAC</option><option>PICC</option><option>Peritoneal tap</option><option>Urinary catheterization</option><option>Lumbar puncture</option><option>Ventricular tap</option><option>PIV</option><option>Chest tube placement</option></select><a href="javascript:void(0);" class="remove_procedure_button" title="Remove field"><i class="fas fa-minus-circle" style="font-size:22px; margin:15px 0 0 5px;"></i></a></div>';
    var y = 1; 
    

    $(add_procedure_button).click(function(){
        if(y < maxField){ 
            y++; 
            $(procedure_wrapper).append(procedureHTML); 
        }
    });
    

    $(procedure_wrapper).on('click', '.remove_procedure_button', function(e){
        e.preventDefault();
        $(this).parent('div').remove(); 
        y--; 
    });
});
